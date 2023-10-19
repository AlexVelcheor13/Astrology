package repository

import (
	"beteraAstrology"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApodRepository interface {
	InsertApod(c *gin.Context, a *beteraAstrology.ApodModel) (int64, error)
	GetApodRecordByDate(date time.Time) (*beteraAstrology.ApodModel, error)
	GetFullAlbum(c *gin.Context) (*[]beteraAstrology.ApodModel, error)
}

type Repository struct {
	ApodRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{ApodRepository: NewApodPostgres(db)}
}
