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

func List(ctx context.Context, db *bun.DB) ([]database.Token, error) {
	var tokens []database.Token

	err := db.NewSelect().Model(&tokens).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func Delete(ctx context.Context, db *bun.DB, tokenID int64) error {
	_, err := db.NewDelete().Model((*database.Token)(nil)).Where("id = ?", tokenID).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
