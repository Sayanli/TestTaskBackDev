package handler

import (
	"context"

	"github.com/Sayanli/TestTaskBackDev/internal/entity"
	"github.com/gofiber/fiber/v2"
)

type CreateUser struct {
	Guid string `json:"guid"`
}

// @Summary Create user
// @Description Create user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body CreateUser true "create user"
// @Success 200 {object} entity.Token
// @Failure 500 {object} error
// @Router /api/v1/auth/create [post]
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	input := new(CreateUser)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}
	tokens, err := h.services.Auth.CreateUser(context.Background(), input.Guid)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to create user", "data": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User created", "data": tokens})
}

// @Summary Refresh token
// @Description Refresh token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body entity.User true "refresh tokens"
// @Success 200 {object} entity.Token
// @Failure 500 {object} error
// @Router /api/v1/auth/refresh [post]
func (h *Handler) RefreshToken(c *fiber.Ctx) error {
	input := new(entity.User)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err.Error()})
	}
	tokens, err := h.services.Auth.RefreshToken(context.Background(), input.Guid, input.RefreshToken)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Failed to refresh token", "data": err.Error()})
	}

	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Token refreshed", "data": tokens})
}
