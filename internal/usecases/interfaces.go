package usecases

import (
	"context"
	"github.com/AutoDivdik/enterprises-service/internal/domain"
	"github.com/google/uuid"
)

type (
	EnterpriseRepo interface {
		GetAll(context.Context) ([]*domain.Enterprise, error)
		GetByID(context.Context, uuid.UUID) (*domain.Enterprise, error)
		Create(context.Context, *domain.Enterprise) error
		Update(context.Context, *domain.Enterprise) (*domain.Enterprise, error)
	}

	UseCase interface {
		GetEnterprisesList(ctx context.Context) ([]*domain.Enterprise, error)
		GetEnterpriseByID(ctx context.Context, id uuid.UUID) (*domain.Enterprise, error)
	}
)
