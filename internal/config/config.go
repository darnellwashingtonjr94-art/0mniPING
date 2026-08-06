package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type TargetConfig struct {
	Name    string `yaml:"name"`
	Address string `yaml:"address"`
	Layer   string `yaml:"layer"`
	Type    string `yaml:"type"`
}

type Config struct {
	Targets []TargetConfig `yaml:"targets"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
