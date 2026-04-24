package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
)

func init() {
}

type Config struct {
	Daemon Daemon
}

type Daemon struct {
	ConfigPath string `env:"HEALTH_CONFIG"`
}

func MustLoad() Config {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("get user config dir: %w", err))
	}

	cfg := Config{
		Daemon: Daemon{
			ConfigPath: filepath.Join(home, ".config", "gober", "health.hcl"),
		},
	}

	if err := env.Parse(&cfg.Daemon); err != nil {
		panic(fmt.Errorf("parse env: %w", err))
	}

	configPath := flag.String("config", cfg.Daemon.ConfigPath, "path to config file")
	flag.Parse()

	cfg.Daemon.ConfigPath = *configPath

	return cfg
}
