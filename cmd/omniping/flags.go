package main

import (
	"flag"
)

type CLIConfig struct {
	Target   string
	Config   string
	Interval string
	Metrics  string
	JSONLog  string
	Debug    bool
}

func ParseFlags() *CLIConfig {
	cfg := &CLIConfig{}
	flag.StringVar(&cfg.Target, "target", "127.0.0.1", "Target IP or hostname")
	flag.StringVar(&cfg.Config, "config", "", "Path to targets YAML config file")
	flag.StringVar(&cfg.Interval, "interval", "", "Cron-style execution interval (e.g. 5s, 1m)")
	flag.StringVar(&cfg.Metrics, "metrics", "", "Prometheus metrics bind address (e.g. :8080)")
	flag.StringVar(&cfg.JSONLog, "json-log", "", "Path to JSON output log file")
	flag.BoolVar(&cfg.Debug, "debug", false, "Enable verbose debug logging")
	flag.Parse()
	return cfg
}
