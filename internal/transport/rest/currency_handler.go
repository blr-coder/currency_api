package rest

import (
	"currency_api/internal/models"
	"currency_api/internal/repository/repository_interfaces"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type CurrencyHandler struct {
	repository repository_interfaces.CurrencyPairRepositoryI
}

func NewCurrencyHandler(repository repository_interfaces.CurrencyPairRepositoryI) *CurrencyHandler {
	return &CurrencyHandler{repository: repository}
}

func (h *CurrencyHandler) Create(c *fiber.Ctx) error {

	pairInput := new(models.CurrencyPairCreateInput)

	err := c.BodyParser(pairInput)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err = pairInput.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	currencyPair, err := h.repository.Create(c.Context(), pairInput)
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	return c.JSON(currencyPair)
}

func (h *CurrencyHandler) List(c *fiber.Ctx) error {

	pairs, err := h.repository.List(c.Context())
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	if pairs != nil {
		fmt.Println("PAIRS", pairs)
	}

	return c.JSON(pairs)
}

func (h *CurrencyHandler) Exchange(c *fiber.Ctx) error {

	exchangeInput := new(models.CurrencyPairExchangeInput)

	err := c.BodyParser(exchangeInput)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	// TODO: Add Validate
	/*if err = exchangeInput.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}*/

	pair, err := h.repository.Get(c.Context(), exchangeInput.CurrencyFrom, exchangeInput.CurrencyTo)
	if err != nil {
		// TODO: Add http status codes
		return err
	}

	return c.JSON(fiber.Map{
		"result": fmt.Sprintf("%.2f %s", exchangeInput.Value*pair.Well, exchangeInput.CurrencyTo),
	})
}
