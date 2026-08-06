package telemetry

import (
	"time"
)

type Summary struct {
	TotalProbes  int
	SuccessCount int
	FailedCount  int
	MinLatency   time.Duration
	MaxLatency   time.Duration
	AvgLatency   time.Duration
}

func Aggregate(results []Result) Summary {
	var sum Summary
	sum.TotalProbes = len(results)
	if sum.TotalProbes == 0 {
		return sum
	}

	var totalLatency time.Duration
	sum.MinLatency = results[0].Latency
	sum.MaxLatency = results[0].Latency

	for _, r := range results {
		if r.Status == StatusUp {
			sum.SuccessCount++
		} else {
			sum.FailedCount++
		}

		totalLatency += r.Latency
		if r.Latency < sum.MinLatency {
			sum.MinLatency = r.Latency
		}
		if r.Latency > sum.MaxLatency {
			sum.MaxLatency = r.Latency
		}
	}
	sum.AvgLatency = totalLatency / time.Duration(sum.TotalProbes)
	return sum
}
