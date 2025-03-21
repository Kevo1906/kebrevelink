package repository

import (
	"context"
	"errors"
	"kebrevelink/internal/domain/models"
	"kebrevelink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormURLRepository struct {
	db *gorm.DB
}

func NewGormURLRepository(db *gorm.DB) repository.URLRepository {
	return &GormURLRepository{db: db}
}

func (r *GormURLRepository) Create(ctx context.Context, url *models.URL) error {
	return r.db.WithContext(ctx).Create(url).Error
}

func (r *GormURLRepository) GetByShortURL(ctx context.Context, shortURL string) (*models.URL, error){
	var url models.URL

	if err := r.db.WithContext(ctx).Where("short_url = ?", shortURL).First(&url).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, err
	}

	return &url,nil
}

func (r *GormURLRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]models.URL, error){
	var urls []models.URL

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&urls).Error; err != nil {		
		return nil, err
	}

	return urls,nil
}

func (r *GormURLRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Where("id = ?",id).Delete(&models.URL{}).Error
}