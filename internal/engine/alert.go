package engine

import (
	"log"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

type AlertManager struct {
	consecutiveFailures map[string]int
	threshold           int
}

func NewAlertManager(threshold int) *AlertManager {
	return &AlertManager{
		consecutiveFailures: make(map[string]int),
		threshold:           threshold,
	}
}

func (a *AlertManager) Evaluate(results []telemetry.Result) {
	for _, res := range results {
		if res.Status != telemetry.StatusUp {
			a.consecutiveFailures[res.Target]++
			if a.consecutiveFailures[res.Target] >= a.threshold {
				log.Printf("[ALERT] Target %s (%s - %s) has failed %d consecutive times! Last error: %s",
					res.Target, res.Layer, res.Probe, a.consecutiveFailures[res.Target], res.Error)
			}
		} else {
			if a.consecutiveFailures[res.Target] >= a.threshold {
				log.Printf("[RESOLVED] Target %s has recovered.", res.Target)
			}
			a.consecutiveFailures[res.Target] = 0
		}
	}
}
