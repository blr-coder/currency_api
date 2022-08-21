package rest

import (
	"currency_api/internal/app/currency/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CreateCurrencyRequest struct {
}

func (h *Handler) Create(c *fiber.Ctx) error {

	pairInput := new(models.CurrencyPairCreateInput)

	err := c.BodyParser(pairInput)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = pairInput.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	currencyPair, err := h.service.Create(c.Context(), pairInput)
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	return c.JSON(currencyPair)
}

func (h *Handler) List(c *fiber.Ctx) error {

	pairs, err := h.service.List(c.Context())
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	return c.JSON(pairs)
}

func (h *Handler) Exchange(c *fiber.Ctx) error {

	exchangeInput := new(models.CurrencyPairExchangeInput)

	err := c.BodyParser(exchangeInput)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// TODO: Add Validate
	/*if err = exchangeInput.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}*/

	pair, err := h.service.Get(c.Context(), exchangeInput.CurrencyFrom, exchangeInput.CurrencyTo)
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	return c.JSON(fiber.Map{
		"result": fmt.Sprintf("%.2f %s", exchangeInput.Value*pair.Well, exchangeInput.CurrencyTo),
	})
}
