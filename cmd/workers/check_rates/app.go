package main

import (
	"context"
	"currency_api/internal/app/config"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/service"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type Worker struct {
	timeout time.Duration
}

func NewWorker(timeout time.Duration) *Worker {
	return &Worker{timeout: timeout}
}

func (w Worker) Run(f func(ctx context.Context) error) {
	logrus.Info("RUN WORKER...")

	ticker := time.NewTicker(w.timeout)
	defer ticker.Stop()

	for range ticker.C {
		err := f(context.TODO())
		if err != nil {
			logrus.Error("RUN WORKER:", err)
		}
	}
}

func runCheckRates() error {

	appConfig, err := config.NewConfig("configs/dev_config.yaml")
	if err != nil {
		return err
	}

	db, err := sqlx.Open("postgres", appConfig.PostgresConnLink)
	if err != nil {
		return err
	}

	r := repository.New(db)
	s := service.New(r, appConfig.AbstractApiKey)

	w := NewWorker(time.Second * 5)
	w.Run(s.CheckRates)

	return nil
}
