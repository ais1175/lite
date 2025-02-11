package token

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
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

func (s *Service) CreateToken(ctx context.Context, data *api.CreateTokenRequest) (string, error) {
	var err error

	apiToken, err := crypt.GenerateApiKey()
	if err != nil {
		return "", nil
	}

	tokenHash, err := crypt.HashPassword(apiToken)
	if err != nil {
		return "", nil
	}

	token := &database.Token{
		Identifier: data.Identifier,
		TokenHash:  tokenHash,
	}

	fmt.Println(token)

	return apiToken, nil
}
