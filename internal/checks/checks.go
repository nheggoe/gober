package checks

import (
	"log/slog"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/nheggoe/gober/internal/config"
)

func Load(cfg config.Daemon) Checks {
	var f HealthConfig
	if err := hclsimple.DecodeFile(cfg.ConfigPath, nil, &f); err != nil {
		slog.Error("failed to load check config:", "err", err)
		os.Exit(1)
	}

	model, err := f.ToRuntimeConfig()
	if err != nil {
		slog.Error("failed to convert config file:", "err", err)
		os.Exit(1)
	}

	return model.Checks
}
