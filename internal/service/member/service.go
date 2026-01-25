package member

import (
	"context"

	"github.com/fivemanage/lite/api"
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

func (s *Service) ListMembers(ctx context.Context, organizationID string) ([]*api.OrganizationMember, error) {
	dbMembers, err := organizationquery.ListMembers(ctx, s.db, organizationID)
	if err != nil {
		return nil, err
	}

	members := make([]*api.OrganizationMember, 0, len(dbMembers))
	for _, dbMember := range dbMembers {
		member := &api.OrganizationMember{
			ID:   dbMember.ID,
			Role: dbMember.Role,
		}

		if dbMember.User != nil {
			member.Email = dbMember.User.Email
			member.Name = dbMember.User.Name
		}

		members = append(members, member)
	}

	return members, nil
}
