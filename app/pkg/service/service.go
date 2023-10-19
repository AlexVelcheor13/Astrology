package service

import (
	"beteraAstrology"
	"beteraAstrology/app/pkg/repository"
	"github.com/gin-gonic/gin"
	"time"
)

type ImageService interface {
	InsertApod(c *gin.Context, a *beteraAstrology.ApodModel) (int64, error)
	GetApodRecordByDate(date time.Time) (*beteraAstrology.ApodModel, error)
	GetFullAlbum(c *gin.Context) (*[]beteraAstrology.ApodModel, error)
}

type Service struct {
	ImageService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{ImageService: NewApodService(repo.ApodRepository)}
}
