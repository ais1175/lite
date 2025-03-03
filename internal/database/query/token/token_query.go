package tokenquery

import (
	"context"

	"github.com/fivemanage/lite/internal/database"
	"github.com/uptrace/bun"
)

func Create(ctx context.Context, db *bun.DB, token *database.Token) error {
	_, err := db.NewInsert().Model(token).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
