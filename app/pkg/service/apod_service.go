package service

import (
	"beteraAstrology"
	"beteraAstrology/app/pkg/repository"
	"github.com/gin-gonic/gin"
	"time"
)

type ApodService struct {
	repo repository.ApodRepository
}

func NewApodService(repo repository.ApodRepository) *ApodService {
	return &ApodService{repo: repo}
}

func (a ApodService) InsertApod(c *gin.Context, am *beteraAstrology.ApodModel) (int64, error) {
	return a.repo.InsertApod(c, am)
}

func (a ApodService) GetApodRecordByDate(date time.Time) (*beteraAstrology.ApodModel, error) {
	return a.repo.GetApodRecordByDate(date)
}

func (a ApodService) GetFullAlbum(c *gin.Context) (*[]beteraAstrology.ApodModel, error) {
	return a.repo.GetFullAlbum(c)
}
