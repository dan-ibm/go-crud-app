package app

import (
	"context"
	http_v1 "github.com/dan-ibm/go-crud-app/internal/delivery/http"
	"github.com/dan-ibm/go-crud-app/internal/repository"
	"github.com/dan-ibm/go-crud-app/internal/service"
	"github.com/dan-ibm/go-crud-app/pkg/database"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func Run() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	// init db
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	db, err := database.NewPostgresConnection(database.ConnectionInfo{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSL"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	// init deps
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := http_v1.NewHandler(services)

	// init & run server
	srv := new(Server)
	go func() {
		if err := srv.Run(os.Getenv("APP_PORT"), handlers.InitRouter()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("CrudApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("CrudApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}
