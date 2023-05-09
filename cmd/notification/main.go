package main

import (
	"context"

	"github.com/caarlos0/env/v6"
	"github.com/khusainnov/notification/internal/app"
	"github.com/khusainnov/notification/internal/config"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()
	cfg := &config.Config{
		L: log,
	}

	if err := env.Parse(cfg); err != nil {
		cfg.L.Fatal("cannot parse config", zap.Error(err))
	}

	if err := app.Run(context.Background(), cfg); err != nil {
		cfg.L.Error("error due run the server", zap.Error(err))
	}
}
