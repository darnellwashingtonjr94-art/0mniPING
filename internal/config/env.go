package config

import (
	"os"
)

func BindEnvOverrides(cfg *Config) {
	if target := os.Getenv("OMNIPING_TARGET"); target != "" {
		cfg.Targets = append(cfg.Targets, TargetConfig{
			Name:    "env-target",
			Address: target,
			Layer:   "L3",
			Type:    "icmp",
		})
	}
}
