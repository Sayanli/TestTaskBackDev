package handler

import (
	"github.com/Sayanli/TestTaskBackDev/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
