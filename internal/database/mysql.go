package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bunotel"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type MySQL struct{}

func (r *MySQL) Connect(dsn string) *bun.DB {
	var sqldb *sql.DB
	var err error

	maxRetries := 3
	retryWaitTime := 5 * time.Second

	for i := range maxRetries {
		sqldb, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("failed to prepare mysql connection: %v", err)
		}

		err = sqldb.Ping()
		if err == nil {
			log.Println("Successfully connected to the database!")
			break
		}

		otelzap.L().Warn("could not connect to the database", zap.Error(err))
		otelzap.S().Warn("Could not connect to the database. Retrying in %v... (Attempt %d/%d)", retryWaitTime, i+1, maxRetries)

		time.Sleep(retryWaitTime)

		if i == maxRetries-1 {
			otelzap.L().Fatal("Failed to connect to the database after multiple attempts", zap.Error(err))
		}
	}

	db := bun.NewDB(sqldb, mysqldialect.New())
	db.AddQueryHook(bunotel.NewQueryHook(bunotel.WithDBName("lite-db")))

	return db
}
