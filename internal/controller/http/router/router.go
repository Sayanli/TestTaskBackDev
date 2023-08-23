package router

import (
	"github.com/Sayanli/TestTaskBackDev/internal/controller/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	app     *fiber.App
	handler *handler.Handler
}

func NewServer(app *fiber.App, handler *handler.Handler) *Server {
	return &Server{app: app, handler: handler}
}

func (s *Server) Router() {

	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	user := v1.Group("/user")
	user.Post("/create", s.handler.CreateUser)
	user.Post("/refresh", s.handler.RefreshToken)
}
