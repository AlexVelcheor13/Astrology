package service

import (
	"beteraAstrology/app"
	"beteraAstrology/app/pkg/repository"
	"github.com/gin-gonic/gin"
	"time"
)

type ImageService interface {
	InsertApod(c *gin.Context, a *app.ApodModel) (int64, error)
	GetApodRecordByDate(date time.Time) (*app.ApodModel, error)
	GetFullAlbum(c *gin.Context) (*[]app.ApodModel, error)
}

type Service struct {
	ImageService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{ImageService: NewApodService(repo.ApodRepository)}
}
