package config

import (
	"errors"
	"fmt"
)

func ValidateConfig(cfg *Config) error {
	if cfg == nil {
		return errors.New("config is nil")
	}
	for i, t := range cfg.Targets {
		if t.Address == "" {
			return fmt.Errorf("target at index %d missing address", i)
		}
		if t.Layer == "" {
			return fmt.Errorf("target %s missing layer definition", t.Name)
		}
	}
	return nil
}
