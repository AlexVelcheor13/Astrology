package service

import (
	"beteraAstrology/app"
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

func (a ApodService) InsertApod(c *gin.Context, am *app.ApodModel) (int64, error) {
	return a.repo.InsertApod(c, am)
}

func (a ApodService) GetApodRecordByDate(date time.Time) (*app.ApodModel, error) {
	return a.repo.GetApodRecordByDate(date)
}

func (a ApodService) GetFullAlbum(c *gin.Context) (*[]app.ApodModel, error) {
	return a.repo.GetFullAlbum(c)
}
