package filequery

import (
	"context"

	"github.com/fivemanage/lite/internal/database"
	"github.com/uptrace/bun"
)

func Create(ctx context.Context, db *bun.DB, file *database.File) (bun.Tx, error) {
	tx, err := db.BeginTx(ctx, nil)
	_, err = tx.NewInsert().Model(file).Exec(ctx)
	if err != nil {
		return tx, err
	}

	return tx, nil
}
