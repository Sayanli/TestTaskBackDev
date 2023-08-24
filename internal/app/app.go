package app

import (
	"github.com/Sayanli/TestTaskBackDev/internal/config"
	"github.com/Sayanli/TestTaskBackDev/internal/controller/http/handler"
	"github.com/Sayanli/TestTaskBackDev/internal/controller/http/router"
	"github.com/Sayanli/TestTaskBackDev/internal/repository"
	"github.com/Sayanli/TestTaskBackDev/internal/service"
	"github.com/Sayanli/TestTaskBackDev/pkg/database/mongodb"
	"github.com/gofiber/fiber/v2"
)

// @title Auth service API
// @version 1.0
// @description This is auth service
// @host 127.0.0.1:8080
// @BasePath /
func Run() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	mongoClient, err := mongodb.NewClient(cfg.MongoDB.URL)
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	repo := repository.NewRepository(mongoClient.Database(cfg.MongoDB.Database))
	service := service.NewService(repo, cfg.JWTSecret.Secret)
	handler := handler.NewHandler(service)
	server := router.NewServer(app, handler)

	server.Router()
	app.Listen(":" + cfg.HTTP.Port)
}
