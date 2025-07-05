package main

import (
	"context"
	"fmt"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fivemanage/lite/internal/clickhouse"
	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/http"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/fivemanage/lite/internal/service/log"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/migrate"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/fivemanage/lite/pkg/kafkaqueue"
	"github.com/fivemanage/lite/pkg/logger"
	"github.com/fivemanage/lite/pkg/otel"
	"github.com/fivemanage/lite/pkg/storage"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:   "fivemanage",
	Short: "Open-source, easy-to-use gaming-community management service.",
	Run: func(cmd *cobra.Command, args []string) {
		var err error

		port := viper.GetInt("port")
		driver := viper.GetString("driver")
		dsn := viper.GetString("dsn")

		sugaredLogger := logger.NewZap()

		defer func() {
			err := sugaredLogger.Sync()
			if err != nil {
				sugaredLogger.Errorf("failed to sync logger: %v", err)
			}
		}()

		otelzap.L().Info("starting Fivemanage application")

		otelShutdown, err := otel.SetupTracer()
		if err != nil {
			sugaredLogger.Fatalf("failed to setup OpenTelemetry: %v", err)
		}

		defer func() {
			sugaredLogger.Info("Attempting to shutdown OpenTelemetry...")
			otelCtx, otelCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer otelCancel()

			if shutdownErr := otelShutdown(otelCtx); shutdownErr != nil {
				sugaredLogger.Errorf("failed to shutdown OpenTelemetry: %v", shutdownErr)
			} else {
				sugaredLogger.Info("OpenTelemetry shutdown successfully.")
			}
		}()

		// this is kinda confusion, 'db' and 'store'
		// maybe we shold call it like...connection and instance?
		// store makes no fucking sense atleast
		db := database.New(driver)
		store := db.Connect(dsn)

		// run migrations for postgres
		// and yes, it's ok to init this here
		migrate.InitMigration(cmd.Context(), store)
		migrate.AutoMigrate(cmd.Context(), store)

		// setup and run migrations for clickhouse
		chConfig := &clickhouse.Config{
			Host:     viper.GetString("clickhouse-host"),
			Username: viper.GetString("clickhouse-username"),
			Password: viper.GetString("clickhouse-password"),
			Database: viper.GetString("clickhouse-database"),
		}

		clickhouse.AutoMigrate(cmd.Context(), chConfig)

		clickhouseClient := clickhouse.NewClient(chConfig)

		kafkaC := kafkaqueue.NewConsumer(otelzap.L())
		kafkaP := kafkaqueue.NewProducer(otelzap.L())

		storageLayer := storage.New("s3")

		authService := auth.NewService(store)
		tokenService := token.NewService(store)
		fileService := file.NewService(store, storageLayer)
		organizationService := organization.NewService(store)
		logService := log.NewService(store, kafkaP, clickhouseClient)

		worker := kafkaqueue.NewBatchWorker(clickhouseClient, kafkaC)

		// kinda not sure what I want to do with this
		// at some point we might actually want to run more than one worker
		go worker.ProcessMessages()

		memcache := cache.NewMemcache(0)

		server := http.NewServer(
			authService,
			tokenService,
			fileService,
			organizationService,
			logService,
			memcache,
		)

		// todo: check if we have an admin user
		// if not, create an admin user with the ADMIN_PASSWORD ENV
		err = authService.CreateAdminUser()
		if err != nil {
			sugaredLogger.Error("failed to create admin user", zap.Error(err))
			return
		}

		srv := &nethttp.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: server,
		}

		go func() {
			otelzap.S().Infof("Fivemanage is running on port %d", port)
			if err := srv.ListenAndServe(); err != nil && err != nethttp.ErrServerClosed {
				sugaredLogger.Fatalf("listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		sugaredLogger.Info("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			sugaredLogger.Fatal("Server Shutdown:", err)
		}

		<-ctx.Done()
		sugaredLogger.Info("timeout of 5 seconds.")
		sugaredLogger.Info("Server exiting")
	},
}

func init() {
	logrus.Println("Initializing Fivemanage...")

	err := godotenv.Load()
	if err != nil {
		logrus.Println("Error loading .env file. Probably becasue we're in production")
	}

	rootCmd.PersistentFlags().String("driver", "pg", "Database driver")
	rootCmd.Flags().Int("port", 8080, "Port to serve Fivemanage")
	rootCmd.Flags().String("dsn", "", "Database DSN")
	rootCmd.Flags().String("clickhouse-host", "localhost:9000", "Clickhouse host")
	rootCmd.Flags().String("clickhouse-username", "default", "Clickhouse username")
	rootCmd.Flags().String("clickhouse-password", "password", "Clickhouse password")
	rootCmd.Flags().String("clickhouse-database", "default", "Clickhouse database")

	// fuck me
	if err := viper.BindPFlag("port", rootCmd.Flags().Lookup("port")); err != nil {
		bindError(err)
	}
	if err := viper.BindPFlag("dsn", rootCmd.Flags().Lookup("dsn")); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("driver", "DB_DRIVER"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("port", "PORT"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("dsn", "DSN"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("clickhouse-host", "CLICKHOUSE_HOST"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("clickhouse-username", "CLICKHOUSE_USERNAME"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("clickhouse-password", "CLICKHOUSE_PASSWORD"); err != nil {
		bindError(err)
	}
	if err := viper.BindEnv("clickhouse-database", "CLICKHOUSE_DATABASE"); err != nil {
		bindError(err)
	}

	rootCmd.AddCommand(migrate.RootCmd)
	migrate.RootCmd.AddCommand(
		migrate.InitCmd,
		migrate.MigrateCmd,
		migrate.CreateMigrationCmd,
		migrate.UnlockCmd,
		migrate.LockCmd,
	)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func bindError(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
