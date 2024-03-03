package usecases

import (
	"context"
	"fmt"
	"github.com/AutoDivdik/enterprises-service/internal/domain"
	"github.com/google/uuid"
)

type usecase struct {
	enterpriseRepo EnterpriseRepo
}

func NewUseCase(enterpriseRepo EnterpriseRepo) UseCase {
	return &usecase{
		enterpriseRepo: enterpriseRepo,
	}
}

func (u *usecase) GetEnterprisesList(ctx context.Context) ([]*domain.Enterprise, error) {
	entities, err := u.enterpriseRepo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("enterpriseRepo.GetAll: %w", err)
	}

	return entities, nil
}

func (u *usecase) GetEnterpriseByID(ctx context.Context, id uuid.UUID) (*domain.Enterprise, error) {
	entity, err := u.enterpriseRepo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("enterpriseRepo.GetEnterpriseByID: %w", err)
	}

	return entity, nil
}
