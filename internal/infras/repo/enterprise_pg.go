package repo

import (
	"context"
	"fmt"
	"github.com/AutoDivdik/enterprises-service/internal/domain"
	postgresql "github.com/AutoDivdik/enterprises-service/internal/infras/pgsql"
	"github.com/AutoDivdik/enterprises-service/internal/usecases"
	engine "github.com/AutoDivdik/pg-engine"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type enterpriseRepo struct {
	pg engine.DBEngine
}

func NewEnterpriseRepo(pg engine.DBEngine) usecases.EnterpriseRepo {
	return &enterpriseRepo{
		pg: pg,
	}
}

func (e *enterpriseRepo) GetAll(ctx context.Context) ([]*domain.Enterprise, error) {
	querier := postgresql.New(e.pg.GetDB())
	results, err := querier.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	enterprises := lo.Map(results, func(x postgresql.GetAllRow, _ int) *domain.Enterprise {
		return &domain.Enterprise{
			ID:              x.ID,
			Name:            x.Name,
			Country:         x.Country,
			MaintenanceYear: int(x.Maintenanceyear),
			Phone:           x.Phone,
			Fax:             x.Fax,
			TypeOfOwnership: &domain.TypeOfOwnership{
				ID:   x.TypeOfOwnershipID,
				Name: x.TypeOfOwnershipName,
			},
		}
	})

	entities := make([]*domain.Enterprise, 0, len(results))

	for _, o := range enterprises {
		enterprise := &domain.Enterprise{
			ID:              o.ID,
			Name:            o.Name,
			Country:         o.Country,
			MaintenanceYear: o.MaintenanceYear,
			Phone:           o.Phone,
			Fax:             o.Fax,
			TypeOfOwnership: &domain.TypeOfOwnership{
				ID:   o.TypeOfOwnership.ID,
				Name: o.TypeOfOwnership.Name,
			},
		}

		entities = append(entities, enterprise)
	}

	return entities, nil
}

func (e *enterpriseRepo) GetByID(ctx context.Context, uuid uuid.UUID) (*domain.Enterprise, error) {
	querier := postgresql.New(e.pg.GetDB())
	row, err := querier.GetEnterpriseByID(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}

	return &domain.Enterprise{
		ID:              row.ID,
		Name:            row.Name,
		Country:         row.Country,
		MaintenanceYear: int(row.Maintenanceyear),
		Phone:           row.Phone,
		Fax:             row.Fax,
		TypeOfOwnership: &domain.TypeOfOwnership{
			ID:   row.TypeOfOwnershipID,
			Name: row.TypeOfOwnershipName,
		},
	}, nil
}

func (e *enterpriseRepo) Create(ctx context.Context, enterprise *domain.Enterprise) error {
	//TODO implement me
	panic("implement me")
}

func (e *enterpriseRepo) Update(ctx context.Context, enterprise *domain.Enterprise) (*domain.Enterprise, error) {
	//TODO implement me
	panic("implement me")
}
