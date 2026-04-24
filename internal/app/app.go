package app

import (
	"context"
	"fmt"

	"github.com/nheggoe/gober/internal/checks"
	"github.com/nheggoe/gober/internal/client"
	"github.com/nheggoe/gober/internal/config"
)

type App struct {
	checks checks.Checks
	client client.Client
}

func New(cfg config.Config) App {
	return App{
		checks: checks.Load(cfg.Daemon),
	}
}

func (a App) Run(ctx context.Context) error {
	for c := range a.checks.All() {
		fmt.Printf("%+v\n", c)
	}
	return nil
}
