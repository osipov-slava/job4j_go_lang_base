package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type DeleteItemRequest struct {
	ID string `json:"id"`
}

func (s *Server) DeleteItem(c *fiber.Ctx) error {
	var req DeleteItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.ID == "" {
		return fiber.NewError(fiber.StatusBadRequest, "id is required")
	}
	err := s.Repository.Delete(c.Context(), req.ID)
	if err != nil {
		// Проверяем тип ошибки
		if strings.Contains(err.Error(), "not found") {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}

		// Логируем внутреннюю ошибку
		log.Errorw("failed to delete item", "id", req.ID, "error", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
