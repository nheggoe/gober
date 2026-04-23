package app

import (
	"context"

	"github.com/nheggoe/gober/internal/checks"
)

type App struct {
	checks checks.Checks
}

func (a App) Run(ctx context.Context) error {
	for c := range a.checks.All() {

	}
	return nil
}
