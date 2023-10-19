package repository

import (
	"beteraAstrology/app"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

type ApodRepository interface {
	InsertApod(c *gin.Context, a *app.ApodModel) (int64, error)
	GetApodRecordByDate(date time.Time) (*app.ApodModel, error)
	GetFullAlbum(c *gin.Context) (*[]app.ApodModel, error)
}

type Repository struct {
	ApodRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{ApodRepository: NewApodPostgres(db)}
}
