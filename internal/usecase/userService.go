package usecase

import (
	"context"
	"errors"
	"kebrevelink/internal/domain/models"
	"kebrevelink/internal/domain/repository"
	"kebrevelink/internal/domain/helpers"
	"strings"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) RegisterUser(ctx context.Context, email string, password string) (*models.User,error){
	email = strings.TrimSpace(email)
	if email == "" || password == ""{
		return nil, errors.New("email and password are required")
	}
	password, err := helpers.HashPassword(password)
	if err != nil{
		return nil, errors.New("the password couldn't be encrypted")
	}

	user := &models.User{
		Email: email,
		Password: password,
	}

	err = s.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error){
	return s.userRepo.GetByEmail(ctx, email)
}