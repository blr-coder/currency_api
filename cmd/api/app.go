package main

import (
	"context"
	"currency_api/internal/app/config"
	"currency_api/internal/app/currency/repository"
	"currency_api/internal/app/currency/service"
	"currency_api/internal/app/currency/transport/rest"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"go.uber.org/multierr"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
)

func init() {
	logrus.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuration
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" -> %s:%d", formatFilePath(f.File), f.Line)
		},
	}
	logrus.SetFormatter(formatter)
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}

func runApp() (err error) {
	logrus.Info("RUN APP...")

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
	s := service.New(r, appConfig.AbstractApiKey)
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

	<-ctx.Done()

	return err
}
