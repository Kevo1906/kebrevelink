package repository

import (
	"context"
	"kebrevelink/internal/domain/models"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error

	GetByID(ctx context.Context, id uuid.UUID) (*models.User, error)

	GetByEmail(ctx context.Context, email string) (*models.User, error)
}