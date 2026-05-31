package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/item/", s.CreateItem)
	route.Get("/items/", s.GetItems)
}
