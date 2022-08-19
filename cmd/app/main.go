package main

import (
	"github.com/dan-ibm/go-crud-app/internal/app"
	"github.com/sirupsen/logrus"
)

// @title CRUD App API
// @version 1.0
// @description API Server for CRUD books Application

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	app.Run()
}
