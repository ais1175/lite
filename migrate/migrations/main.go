package migrations

import (
	"embed"

	"github.com/uptrace/bun/migrate"
)

//go:embed *.go
var sqlMigrations embed.FS

var Migrations = migrate.NewMigrations()

func init() {
	if err := Migrations.Discover(sqlMigrations); err != nil {
		panic(err)
	}
}
