package repository

import (
	"beteraAstrology/app"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	queryInsert = `INSERT INTO apod
				   (date, copyright, explanation, hd_url, media_type, service_version, title, url) 
				   VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	queryGetByDate = `SELECT date, copyright, explanation, hd_url, media_type, service_version, title, url
					  FROM apod WHERE "date" = $1`

	queryGetFullAlbum = `SELECT date, copyright, explanation, hd_url, media_type, service_version, title, url
					 	   FROM apod`
)

type ApodPostgres struct {
	db *sqlx.DB
}

func NewApodPostgres(db *sqlx.DB) *ApodPostgres {
	return &ApodPostgres{db: db}
}

func (a ApodPostgres) InsertApod(c *gin.Context, am *app.ApodModel) (int64, error) {

	result, err := a.db.ExecContext(c, queryInsert, am.Date, am.Copyright, am.Explanation, am.HDURL, am.MediaType, am.ServiceVersion,
		am.Title, am.URL)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (a ApodPostgres) GetApodRecordByDate(date time.Time) (*app.ApodModel, error) {
	var apod app.ApodModel

	if err := a.db.Get(&apod, queryGetByDate, date.Format("2006-01-02")); err != nil {
		return nil, err
	}
	if &apod == nil {
		return nil, errors.New("nothing to show")
	}

	return &apod, nil
}

func (a ApodPostgres) GetFullAlbum(c *gin.Context) (*[]app.ApodModel, error) {
	var apods []app.ApodModel

	if err := a.db.SelectContext(c, &apods, queryGetFullAlbum); err != nil {
		return nil, err
	}
	if apods == nil {
		return nil, errors.New("nothing to show")
	}

	return &apods, nil
}
