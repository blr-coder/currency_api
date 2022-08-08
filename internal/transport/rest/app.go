package rest

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

type CurrencyApp struct {
	app             *fiber.App
	currencyHandler *CurrencyHandler
}

func NewCurrencyApp(currencyHandler *CurrencyHandler) *CurrencyApp {
	return &CurrencyApp{currencyHandler: currencyHandler}
}

func (a *CurrencyApp) Run(port string) error {

	app := fiber.New(fiber.Config{
		AppName:      "CURRENCY_API",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	})

	app.Post("/api/currency", a.currencyHandler.Create)
	app.Put("/api/currency", a.currencyHandler.Exchange)
	app.Get("/api/currency", a.currencyHandler.List)

	return app.Listen(fmt.Sprintf(":%s", port))
}
