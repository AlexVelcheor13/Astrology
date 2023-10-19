package main

import (
	withGin "beteraAstrology"
	"beteraAstrology/app/pkg/handler"
	repository2 "beteraAstrology/app/pkg/repository"
	"beteraAstrology/app/pkg/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	startServer()
}

func startServer() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository2.NewPostgresDB(repository2.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("faild to init db: %s", err.Error())
	}

	repos := repository2.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	serv := new(withGin.Server)
	go func() {
		if err := serv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error to run server: %s", err.Error())
		}
	}()
	logrus.Print("Api started")

	c := gin.Context{}
	scheduler := gocron.NewScheduler(time.UTC)
	_, err = scheduler.Every(1).Day().At("11:00").Do(handlers.InsertApod, &c)
	if err != nil {
		logrus.Fatalf("err ocured on shaduler: %s", err.Error())
		return
	}

	defer scheduler.Stop()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Astronomy api shutting down")
	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("err ocured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Fatalf("err ocured on db connection close: %s", err.Error())
	}
}
