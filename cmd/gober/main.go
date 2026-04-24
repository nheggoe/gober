package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/nheggoe/gober/internal/app"
	"github.com/nheggoe/gober/internal/config"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	cfg := config.MustLoad()

	if err := app.New(cfg).Run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run",
			slog.String("err", err.Error()),
		)
	}
}
