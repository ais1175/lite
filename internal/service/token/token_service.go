package token

import (
	"context"
	"strconv"

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

func (r *Service) ListTokens(ctx context.Context) ([]api.ListTokensResponse, error) {
	var err error
	var response []api.ListTokensResponse

	tokens, err := tokenquery.List(ctx, r.db)
	if err != nil {
		return nil, err
	}

	for _, token := range tokens {
		response = append(response, api.ListTokensResponse{
			ID:         token.ID,
			Identifier: token.Identifier,
		})
	}

	return response, nil
}

func (r *Service) DeleteToken(ctx context.Context, tokenID string) error {
	var err error
	iTokenID, err := strconv.ParseInt(tokenID, 10, 64)
	if err != nil {
		return err
	}

	err = tokenquery.Delete(ctx, r.db, iTokenID)
	if err != nil {
		return err
	}
	return nil
}
