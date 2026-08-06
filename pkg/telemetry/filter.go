package telemetry

import "time"

func FilterFailed(results []Result) []Result {
	var failed []Result
	for _, r := range results {
		if r.Status != StatusUp {
			failed = append(failed, r)
		}
	}
	return failed
}

func FilterSlow(results []Result, threshold time.Duration) []Result {
	var slow []Result
	for _, r := range results {
		if r.Latency > threshold {
			slow = append(slow, r)
		}
	}
	return slow
}
