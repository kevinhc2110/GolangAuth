package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/kevinhc2110/Auth_UCP/internal/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	FindByID(ctx context.Context, ID uuid.UUID) (*domain.User, error)
	FindByIdentification(ctx context.Context, identification string) (*domain.User, error)
	FindByEmail(ctx context.Context, email string) (*domain.User, error)
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
