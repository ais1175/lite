package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type MySQL struct{}

func (r *MySQL) Connect(dsn string) *bun.DB {
	sqldb, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open mysql connection: %v", err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}
