package api

import "github.com/gofiber/fiber/v2"

func (s *Server) Route(route fiber.Router) {
	route.Post("/items/", s.CreateItem)
	route.Get("/items/", s.GetItems)
	route.Put("/items/", s.UpdateItem)
	route.Delete("/items/", s.DeleteItem)
}
