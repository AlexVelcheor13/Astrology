package handler

import (
	"beteraAstrology/app"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	Date     = "date"
	Token    = "NASA_API_KEY_TOKEN"
	ApiKey   = "?api_key="
	NASA_URL = "NASA_URL"
)

func (h *Handler) InsertApod(c *gin.Context) {
	var apod app.ApodModel

	response, err := http.Get(os.Getenv(NASA_URL) + ApiKey + os.Getenv(Token))
	if err != nil {
		logrus.Warning("something is wrong with response")
		return
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		logrus.Warning("something is wrong with reading response body")
		return
	}

	err = json.Unmarshal(body, &apod)
	if err != nil {
		logrus.Warning("something is wrong with unmarshalling")
		newErrorResponse(c, http.StatusInternalServerError, "error on unmarshalling")
		return
	}

	_, err = h.services.InsertApod(c, &apod)
	if err != nil {
		logrus.Warning("something is wrong in InsertApod")
		return
	}

	qualityURL, err := chooseQualityURL(&apod)
	if err != nil {
		return
	}

	bytes, err := downloadImage(c, qualityURL)
	if err != nil {
		return
	}

	err = saveImageOnDisk("./pictures", &apod, bytes)
	if err != nil {
		return
	}
}

func (h *Handler) GetApodByDate(c *gin.Context) {
	dateStr := c.Params.ByName(Date)
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return
	}

	apodModel, err := h.services.GetApodRecordByDate(date)
	if err != nil {
		logrus.Warning("something is wrong")
		newErrorResponse(c, http.StatusInternalServerError, "error on getting record")
		return
	}

	if &apodModel != nil {
		c.JSON(http.StatusOK, apodModel)
	} else {
		c.JSON(http.StatusInternalServerError, "nothing to get")
	}

}

func (h *Handler) GetFullAlbum(c *gin.Context) {
	fullAlbum, err := h.services.GetFullAlbum(c)
	if err != nil {
		logrus.Warningf("something is wrong %v", err.Error())
		newErrorResponse(c, http.StatusInternalServerError, "error on getting full album")
		return
	}

	c.JSON(http.StatusOK, &fullAlbum)
}

type Response struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, Response{Message: message})
}
