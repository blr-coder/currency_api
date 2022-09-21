package main

import (
	"context"
	"currency_api/internal/app/config"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/service"
	"currency_api/internal/app/currency/transport/rest"
	"currency_api/pkg/exchange_rates"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"go.uber.org/multierr"
	"os/signal"
	"syscall"
	"time"
)

func runApp() (err error) {

	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT)
	defer cancel()

	appConfig, err := config.NewConfig("configs/dev_config.yaml")
	if err != nil {
		return err
	}

	db, err := sqlx.Open("postgres", appConfig.PostgresConnLink)
	if err != nil {
		return err
	}

	r := repository.New(db)
	s := service.New(r)
	api := fiber.New()

	h := rest.New(s, api)
	h.InitCurrencyRoutes()

	// API
	go func() {
		listenErr := api.Listen(fmt.Sprintf(":%s", appConfig.Port))
		if listenErr != nil {
			multierr.AppendInto(&err, listenErr)
			cancel()
			return
		}
	}()

	// Check rates
	abstractClient := exchange_rates.New(appConfig.AbstractApiKey)

	go func() {

		for tick := range time.Tick(20 * time.Second) {
			fmt.Println("Tick", tick.UTC().Format(time.RFC3339))

			listCurrencyPairs, err := r.Pair.List(ctx)
			if err != nil {
				// log errors
				return
			}

			currencyMap := listCurrencyPairs.MapByCurrency()

			for f, t := range currencyMap {
				rates, err := abstractClient.GetRates(ctx, f, t)
				if err != nil {
					// log errors
				}

				_ = r.Pair.UpdateCurrencyWell(ctx, rates)
				if err != nil {
					// log errors
				}

				time.Sleep(2 * time.Second)
			}
		}

	}()

	<-ctx.Done()

	return err
}
