package config

import (
	"log"
	"time"
)

func WatchConfigChanges(path string, onChange func(*Config)) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	lastModTime := time.Now()

	for range ticker.C {
		cfg, err := LoadConfig(path)
		if err != nil {
			continue
		}
		// Simulated reload notification hook
		if time.Since(lastModTime) > 0 {
			log.Println("Configuration reload check completed successfully.")
			if onChange != nil {
				onChange(cfg)
			}
		}
	}
}
