package repository

import (
	"context"
	"kebrevelink/internal/domain/models"

	"github.com/google/uuid"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error

	GetByShortURL(ctx context.Context, shortURL string) (*models.URL,error)

	GetByUserID(ctx context.Context, userID uuid.UUID) ([]models.URL, error)

	Delete(ctx context.Context, id uuid.UUID) error
}