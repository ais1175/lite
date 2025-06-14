package clickhouse

import (
	"context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/fivemanage/lite/pkg/project"
	"github.com/golang-migrate/migrate/v4"
	clickhouseMigrate "github.com/golang-migrate/migrate/v4/database/clickhouse"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func AutoMigrate(ctx context.Context, config *Config) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	options := getClickhouseOptions(config)
	db := clickhouse.OpenDB(options)
	defer db.Close()

	if err := db.PingContext(ctx); err != nil {
		fmt.Println("Error pinging database:", err)
		panic(err)
	}

	driver, err := clickhouseMigrate.WithInstance(db, &clickhouseMigrate.Config{
		MigrationsTableEngine: "MergeTree",
		MultiStatementEnabled: true,
	})
	if err != nil {
		fmt.Println("Error creating ClickHouse driver:", err)
		panic(err)
	}

	migrationsPath := filepath.Join(project.GetRoot(), "internal", "clickhouse", "migrations")
	absPath, err := filepath.Abs(migrationsPath)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		panic(err)
	}

	fmt.Println("Migrations path:", absPath)

	sourcePath := filepath.ToSlash(absPath)
	sourceURL := "file://" + sourcePath

	m, err := migrate.NewWithDatabaseInstance(sourceURL, "default", driver)
	if err != nil {
		fmt.Println("Error creating migration instance:", err)
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error running migrations:", err)
		panic(err)
	}

	fmt.Println("Migrations completed successfully.")
}
