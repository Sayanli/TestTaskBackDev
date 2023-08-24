package router

import (
	_ "github.com/Sayanli/TestTaskBackDev/docs"
	"github.com/Sayanli/TestTaskBackDev/internal/controller/http/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type Server struct {
	app     *fiber.App
	handler *handler.Handler
}

func NewServer(app *fiber.App, handler *handler.Handler) *Server {
	return &Server{app: app, handler: handler}
}

func (s *Server) Router() {

	s.app.Get("/swagger/*", swagger.HandlerDefault)

	api := s.app.Group("/api", logger.New())
	v1 := api.Group("/v1")

	auth := v1.Group("/auth")
	auth.Post("/create", s.handler.CreateUser)
	auth.Post("/refresh", s.handler.RefreshToken)
}
