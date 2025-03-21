package repository

import (
	"context"
	"errors"
	"kebrevelink/internal/domain/models"
	"kebrevelink/internal/domain/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db:db}
}

func (r *GormUserRepository) Create(ctx context.Context, user *models.User) error{
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *GormUserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error){
	var user models.User
	if err:= r.db.WithContext(ctx).First(&user, "id = ?",id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, err
	}
	return &user,nil
}

func (r *GormUserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error){
	var user models.User
	if err:= r.db.WithContext(ctx).First(&user, "email = ?",email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}