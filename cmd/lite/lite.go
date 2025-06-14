package main

import (
	"context"
	"fmt"
	"log"
	nethttp "net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fivemanage/lite/internal/database"
	"github.com/fivemanage/lite/internal/http"
	"github.com/fivemanage/lite/internal/service/auth"
	"github.com/fivemanage/lite/internal/service/file"
	"github.com/fivemanage/lite/internal/service/organization"
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/migrate"
	"github.com/fivemanage/lite/pkg/cache"
	"github.com/fivemanage/lite/pkg/logger"
	"github.com/fivemanage/lite/pkg/otel"
	"github.com/fivemanage/lite/pkg/storage"
	"github.com/joho/godotenv"
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
				log.Printf("failed to sync logger: %v", err)
			}
		}()

		otelzap.L().Info("starting Fivemanage application")

		otelShutdown, err := otel.SetupTracer()
		if err != nil {
			sugaredLogger.Fatalf("failed to setup OpenTelemetry: %v", err)
		}

		defer func() {
			log.Println("Attempting to shutdown OpenTelemetry...")
			otelCtx, otelCancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer otelCancel()

			if shutdownErr := otelShutdown(otelCtx); shutdownErr != nil {
				log.Printf("failed to shutdown OpenTelemetry: %v", shutdownErr)
			} else {
				log.Println("OpenTelemetry shutdown successfully.")
			}
		}()

		db := database.New(driver)
		store := db.Connect(dsn)

		migrate.InitMigration(cmd.Context(), store)
		migrate.AutoMigrate(cmd.Context(), store)

		storageLayer := storage.New("s3")

		authservice := auth.NewService(store)
		tokenservice := token.NewService(store)
		fileservice := file.NewService(store, storageLayer)
		organizationservice := organization.NewService(store)

		memcache := cache.NewMemcache(0)

		server := http.NewServer(
			authservice,
			tokenservice,
			fileservice,
			organizationservice,
			memcache,
		)

		// todo: check if we have an admin user
		// if not, create an admin user with the ADMIN_PASSWORD ENV
		err = authservice.CreateAdminUser()
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
				log.Fatalf("listen: %s\n", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		log.Println("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		select {
		case <-ctx.Done():
			log.Println("timeout of 5 seconds.")
		}
		log.Println("Server exiting")
	},
}

func init() {
	log.Println("Initializing Fivemanage...")

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file. Probably becasue we're in production")
	}

	rootCmd.PersistentFlags().String("driver", "pg", "Database driver")
	rootCmd.Flags().Int("port", 8080, "Port to serve Fivemanage")
	rootCmd.Flags().String("dsn", "", "Database DSN")

	viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
	viper.BindPFlag("dsn", rootCmd.Flags().Lookup("dsn"))
	viper.BindEnv("driver", "DB_DRIVER")
	viper.BindEnv("port", "PORT")
	viper.BindEnv("dsn", "DSN")

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
