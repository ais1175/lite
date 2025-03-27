package organization

import (
	"context"

	"github.com/fivemanage/lite/api"
	"github.com/fivemanage/lite/internal/database"
	organizationquery "github.com/fivemanage/lite/internal/database/query/organization"
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

func (r *Service) CreateOrganization(ctx context.Context, data *api.CreateOrganizationRequest) (*api.Organization, error) {
	dbOrganization := &database.Organization{
		Name: data.Name,
	}

	_, err := organizationquery.Create(ctx, r.db, dbOrganization)
	if err != nil {
		return nil, err
	}

	organization := &api.Organization{
		ID:   dbOrganization.ID,
		Name: dbOrganization.Name,
	}

	return organization, nil
}

func (r *Service) FindOrganizationByID(ctx context.Context, id int64) (*api.Organization, error) {
	dbOrganization, err := organizationquery.Find(ctx, r.db, id)
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
