package handler

import (
	"beteraAstrology/app"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func chooseQualityURL(a *app.ApodModel) (string, error) {
	var urlQuality string
	switch {
	case a.HDURL != "":
		urlQuality = a.HDURL
	case a.URL != "":
		urlQuality = a.URL
	default:
		logrus.Error("Something is wrong in url")
	}

	return urlQuality, nil
}

func downloadImage(c *gin.Context, url string) ([]byte, error) {
	request, err := http.NewRequestWithContext(c, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}

func saveImageOnDisk(path string, a *app.ApodModel, image []byte) error {
	a.FreshRaw = image
	imageName := filepath.Join(path, a.Date+"."+a.Title)

	_, err := os.Stat(imageName)
	if errors.Is(err, os.ErrExist) {
		return nil
	}
	file, err := os.Create(imageName)
	if err != nil {
		return err
	}

	err = file.Chmod(os.FileMode(0777))
	if err != nil {
		return err
	}

	_, err = file.Write(a.FreshRaw)
	if err != nil {
		return err
	}

	return nil
}
