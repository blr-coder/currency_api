package main

import (
	"currency_api/config"
	"currency_api/internal/repository"
	"currency_api/internal/storage/postgres"
	"currency_api/internal/transport/rest"
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {

	appConfig := config.NewConfig()

	psqlDB, err := postgres.NewDB(
		appConfig.Postgres.Host,
		appConfig.Postgres.Port,
		appConfig.Postgres.Name,
		appConfig.Postgres.User,
		appConfig.Postgres.Password,
	)
	if err != nil {
		panic(fmt.Errorf("failed to init postgres: %v", err.Error()))
	}

	logger := logrus.StandardLogger()

	r := repository.NewCurrencyPairRepository(psqlDB, logger)

	h := rest.NewCurrencyHandler(r)

	app := rest.NewCurrencyApp(h)

	err = app.Run(appConfig.Port)
	if err != nil {
		panic(fmt.Errorf("failed to run app: %v", err.Error()))
	}
}
