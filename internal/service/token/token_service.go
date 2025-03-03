package token

import (
	"context"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	tokenquery "github.com/fivemanage/lite/internal/database/query/token"
	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (r *Service) CreateToken(ctx context.Context, data *api.CreateTokenRequest) (string, error) {
	var err error

	apiToken, err := crypt.GenerateApiKey()
	if err != nil {
		return "", err
	}

	tokenHash, err := crypt.HashPassword(apiToken)
	if err != nil {
		return "", err
	}

	token := &database.Token{
		Identifier: data.Identifier,
		TokenHash:  tokenHash,
	}

	err = tokenquery.Create(ctx, r.db, token)
	if err != nil {
		return "", err
	}

	return apiToken, nil
}
