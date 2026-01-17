package organization

import (
	"context"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/clickhouse"
	"github.com/fivemanage/lite/internal/crypt"
	"github.com/fivemanage/lite/internal/database"
	datasetquery "github.com/fivemanage/lite/internal/database/query/dataset"
	filequery "github.com/fivemanage/lite/internal/database/query/file"
	organizationquery "github.com/fivemanage/lite/internal/database/query/organization"
	tokenquery "github.com/fivemanage/lite/internal/database/query/token"
	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
	ch *clickhouse.Client
}

func NewService(db *bun.DB, ch *clickhouse.Client) *Service {
	return &Service{
		db: db,
		ch: ch,
	}
}

func (r *Service) GetStats(ctx context.Context, organizationID string) (*api.OrganizationStats, error) {
	totalFiles, err := filequery.FindTotalStorageCount(ctx, r.db, organizationID)
	if err != nil {
		return nil, err
	}

	totalSize, err := filequery.FindTotalStorageSize(ctx, r.db, organizationID)
	if err != nil {
		return nil, err
	}

	totalLogs, err := r.ch.QueryTotalLogsByOrg(ctx, organizationID)
	if err != nil {
		return nil, err
	}

	tokens, err := tokenquery.List(ctx, r.db, organizationID)
	if err != nil {
		return nil, err
	}

	datasets, err := datasetquery.List(ctx, r.db, organizationID)
	if err != nil {
		return nil, err
	}

	return &api.OrganizationStats{
		TotalFiles:   totalFiles,
		TotalSize:    totalSize,
		TotalLogs:    totalLogs,
		TotalTokens:  len(tokens),
		TotalDataset: len(datasets),
	}, nil
}

func (r *Service) CreateOrganization(ctx context.Context, data *api.CreateOrganizationRequest, userID int64) (*api.Organization, error) {
	orgId, err := crypt.GeneratePrimaryKey()
	if err != nil {
		return nil, err
	}

	dbOrganization := &database.Organization{
		Name: data.Name,
		ID:   orgId,
	}

	orgTx, err := organizationquery.Create(ctx, r.db, dbOrganization)
	if err != nil {
		return nil, err
	}
	if err := orgTx.Commit(); err != nil {
		return nil, err
	}

	memberTx, err := organizationquery.CreateMember(ctx, r.db, &database.OrganizationMember{
		Role:           "ADMIN",
		OrganizationID: dbOrganization.ID,
		UserID:         userID,
	})
	if err != nil {
		return nil, err
	}

	if err := memberTx.Commit(); err != nil {
		if err := orgTx.Rollback(); err != nil {
			return nil, err
		}
		return nil, err
	}

	organization := &api.Organization{
		ID:   dbOrganization.ID,
		Name: dbOrganization.Name,
	}

	return organization, nil
}

func (r *Service) FindOrganizationByID(ctx context.Context, ID string) (*api.Organization, error) {
	dbOrganization, err := organizationquery.Find(ctx, r.db, ID)
	if err != nil {
		return nil, err
	}

	organization := &api.Organization{
		ID:   dbOrganization.ID,
		Name: dbOrganization.Name,
	}

	return organization, nil
}

func (r *Service) ListOrganizations(ctx context.Context) ([]*api.Organization, error) {
	dbOrganizations, err := organizationquery.List(ctx, r.db)
	if err != nil {
		return nil, err
	}

	organizations := make([]*api.Organization, 0, len(dbOrganizations))
	for _, dbOrganization := range dbOrganizations {
		organization := &api.Organization{
			ID:   dbOrganization.ID,
			Name: dbOrganization.Name,
		}

		organizations = append(organizations, organization)
	}

	return organizations, nil
}
