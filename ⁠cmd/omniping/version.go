package main

import (
	"fmt"
)

var (
	Version   = "v1.0.0"
	BuildTime = "unknown"
)

func PrintVersion() {
	fmt.Printf("omniping version %s (built at %s)\n", Version, BuildTime)
}
