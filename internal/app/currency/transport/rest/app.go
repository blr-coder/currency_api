package rest

import (
	"github.com/gofiber/fiber/v2"

	"currency_api/internal/app/currency/service"
)

type Handler struct {
	api     *fiber.App
	service *service.Service
}

func New(service *service.Service, api *fiber.App) {
	h := &Handler{
		service: service,
		api:     api,
	}

	h.api.Post("/api/currency", h.Create)
	h.api.Put("/api/currency", h.Exchange)
	h.api.Get("/api/currency", h.List)
}
