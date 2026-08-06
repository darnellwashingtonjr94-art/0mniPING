package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/username/omniping/internal/engine"
	"github.com/username/omniping/internal/exporter"
)

func main() {
	targetFlag := flag.String("target", "127.0.0.1", "Target IP or hostname to ping")
	flag.Parse()

	ctx := context.Background()
	runner := engine.NewRunner([]string{*targetFlag})
	results := runner.RunAll(ctx)

	exporter.PrintResults(results)

	for _, r := range results {
		if r.Status != telemetry.StatusUp {
			os.Exit(1)
		}
	}
}
