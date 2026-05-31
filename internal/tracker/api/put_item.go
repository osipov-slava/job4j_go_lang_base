package api

import (
	"database/sql"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"job4j.ru/go-lang-base/internal/tracker"
)

func (s *Server) UpdateItem(c *fiber.Ctx) error {
	var req ItemRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid JSON body")
	}
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "name is required")
	}

	item := tracker.Item{req.ID, req.Name}
	err := s.Repository.Update(c.Context(), req.ID, item)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Errorw("item not found", "id", req.ID, "error", err)
			return fiber.NewError(fiber.StatusNotFound, "item not found")
		}
		log.Errorw("failed to update item", "id", req.ID, "error", err)
		return fiber.NewError(fiber.StatusInternalServerError, "internal server error")
	}
	return nil
}
