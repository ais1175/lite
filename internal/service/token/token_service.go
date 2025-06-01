package token

import (
	"context"
	"strconv"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	tokenquery "github.com/fivemanage/lite/internal/database/query/token"
	"github.com/uptrace/bun"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.uber.org/zap"
)

var tracer = otel.Tracer("github.com/fivemanage/lite/internal/service/token")

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (r *Service) GetToken(ctx context.Context, tokenHash string) (*database.Token, error) {
	var err error

	token, err := tokenquery.SelectByHash(ctx, r.db, tokenHash)
	if err != nil {
		return nil, err
	}

	return token, nil
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
		OrganizationID: data.OrganizationID,
		Identifier:     data.Identifier,
		TokenHash:      tokenHash,
	}

	err = tokenquery.Create(ctx, r.db, token)
	if err != nil {
		return "", err
	}

	return apiToken, nil
}

func (r *Service) ListTokens(ctx context.Context) ([]*api.ListTokensResponse, error) {
	var err error
	var response []*api.ListTokensResponse

	ctx, span := tracer.Start(ctx, "token.list_tokens")
	defer span.End()

	tokens, err := tokenquery.List(ctx, r.db)
	if err != nil {
		span.RecordError(err)
		otelzap.L().Error("failed to list tokens", zap.Error(err))
		return nil, err
	}

	for _, token := range tokens {
		response = append(response, &api.ListTokensResponse{
			ID:         token.ID,
			Identifier: token.Identifier,
		})
	}

	span.SetAttributes(
		attribute.Int("token_count", len(tokens)),
	)

	otelzap.L().Debug("listed tokens successfully", zap.Int("count", len(tokens)))

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
