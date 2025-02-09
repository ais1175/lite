package token

import (
	"context"
	"fmt"

	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	"github.com/uptrace/bun"
)

type Service struct{}

func NewService(db *bun.DB) *Service {
	return &Service{}
}

func (s *Service) CreateToken(ctx context.Context, identifier, tokenType string) error {
	var err error

	apiToken, err := crypt.GenerateApiKey()
	if err != nil {
		return err
	}

	tokenHash, err := crypt.HashPassword(apiToken)
	if err != nil {
		return err
	}

	token := &database.Token{
		Identifier: identifier,
		TokenHash:  tokenHash,
	}

	fmt.Println(token)
	return nil
}
