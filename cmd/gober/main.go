package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/nheggoe/gober/internal/app"
	"github.com/nheggoe/gober/internal/buildinfo"
	"github.com/nheggoe/gober/internal/config"
)

func main() {
	showVersion := flag.Bool("version", false, "print version number")
	flag.Parse()

	if *showVersion {
		printVersion()
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGTERM, os.Interrupt)
	defer stop()

	cfg := config.MustLoad()

	if err := app.New(cfg).Run(ctx); err != nil {
		slog.ErrorContext(ctx, "failed to run",
			slog.String("err", err.Error()),
		)
	}
}

func printVersion() {
	fmt.Printf("gober %s\n", buildinfo.Version)
}
