package usecase

import (
	"context"
	"errors"
	"kebrevelink/internal/domain/models"
	"kebrevelink/internal/domain/repository"
	"strings"
	"time"

	"github.com/google/uuid"
)

type URLService struct {
	urlRepo repository.URLRepository
}

func NewURLService(urlRepo repository.URLRepository) *URLService {
	return &URLService{urlRepo: urlRepo}
}

func (s *URLService) CreateShortURL(ctx context.Context, originalURL string, userID uuid.UUID, expiration *time.Time) (*models.URL, error){
	originalURL = strings.TrimSpace(originalURL)
	if originalURL == ""{
		return nil, errors.New("original URL is required")
	}

	url := &models.URL{
		Original: originalURL,
		ShortUrl: generateShortCode(),
		UserID: userID,
		ExpiresAt: expiration,
	}

	err := s.urlRepo.Create(ctx,url)
	if err != nil {
		return nil,err
	}
	return url, nil
}

func (s *URLService) GetOriginalURL(ctx context.Context, shorCode string)(string,error){
	url, err := s.urlRepo.GetByShortURL(ctx,shorCode)
	if err != nil{
		return "", err
	}
	if url == nil || (url.ExpiresAt != nil && url.ExpiresAt.Before(time.Now())){
		return "", errors.New("URL not found or expired")
	}
	return url.Original, nil
}

func generateShortCode() string{
	return ""
}