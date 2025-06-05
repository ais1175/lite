package database

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type PostgreSQL struct{}

func (r *PostgreSQL) Connect(dsn string) *bun.DB {
	var sqldb *sql.DB
	var err error

	maxRetries := 3
	retryWaitTime := 5 * time.Second

	for i := range maxRetries {
		sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
		if err != nil {
			otelzap.L().Fatal("Failed to prepare PostgreSQL connection", zap.Error(err))
		}

		err = sqldb.Ping()
		if err == nil {
			otelzap.L().Info("Successfully connected to the PostgreSQL database")
			break
		}

		otelzap.L().Warn("could not connect to the database", zap.Error(err))
		otelzap.S().Warn("Could not connect to the database. Retrying in %v... (Attempt %d/%d)", retryWaitTime, i+1, maxRetries)

		time.Sleep(retryWaitTime)

		if i == maxRetries-1 {
			otelzap.L().Fatal("Failed to connect to the database after multiple attempts", zap.Error(err))
		}
	}

	return bun.NewDB(sqldb, pgdialect.New())
}
