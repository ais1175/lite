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
	"github.com/fivemanage/lite/internal/service/token"
	"github.com/fivemanage/lite/internal/storage"
	"github.com/fivemanage/lite/migrate"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "fivemanage",
		Short: "Open-source, easy-to-use gaming-community management service.",
	}

	// todo: otel & promhttp

	// this whole code can probably go into a 'run.go' file.
	// just to not fill this file with shit
	// its gonna be ugly anyways tho
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "Run fivemanage application",
		Run: func(cmd *cobra.Command, args []string) {
			var err error

			port := viper.GetInt("port")
			driver := viper.GetString("driver")

			err = godotenv.Load()
			// TODO: Only for development
			if err != nil {
				log.Fatal("Error loading .env file. Probably becasue we're in production")
			}

			db := database.New(driver, "")
			store := db.Connect()
			migrate.AutoMigrate(cmd.Context(), store)

			storageLayer := storage.New("s3")

			authservice := auth.New(store)
			tokenservice := token.NewService(store)
			fileservice := file.NewService(store, storageLayer)

			server := http.NewServer(
				authservice,
				tokenservice,
				fileservice,
			)

			// todo: check if we have an admin user
			// if not, create an admin user with the ADMIN_PASSWORD ENV
			err = authservice.CreateAdminUser()
			if err != nil {
				logrus.WithError(err).Error("Failed to create admin user")
				return
			}

			srv := &nethttp.Server{
				Addr:    fmt.Sprintf("localhost:%d", port),
				Handler: server,
			}

			go func() {
				fmt.Printf("Server is running on port %d...\n", port)
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
)

func init() {
	rootCmd.PersistentFlags().String("driver", "mysql", "Database driver")
	runCmd.Flags().Int("port", 8080, "Port to serve Fivemanage")

	viper.BindPFlag("driver", rootCmd.PersistentFlags().Lookup("driver"))
	viper.BindPFlag("port", runCmd.Flags().Lookup("port"))

	rootCmd.AddCommand(runCmd)

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
