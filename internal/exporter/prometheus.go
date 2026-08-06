package exporter

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/username/omniping/pkg/telemetry"
)

type PrometheusServer struct {
	mu      sync.Mutex
	results []telemetry.Result
	addr    string
}

func NewPrometheusServer(addr string) *PrometheusServer {
	return &PrometheusServer{addr: addr}
}

func (s *PrometheusServer) Update(results []telemetry.Result) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.results = results
}

func (s *PrometheusServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	defer s.mu.Unlock()

	w.Header().Set("Content-Type", "text/plain; version=0.0.4")
	for _, res := range s.results {
		upVal := 0
		if res.Status == telemetry.StatusUp {
			upVal = 1
		}
		fmt.Fprintf(w, "omniping_up{target=\"%s\",layer=\"%s\",probe=\"%s\"} %d\n", res.Target, res.Layer, res.Probe, upVal)
		fmt.Fprintf(w, "omniping_latency_seconds{target=\"%s\",layer=\"%s\",probe=\"%s\"} %.5f\n", res.Target, res.Layer, res.Probe, res.Latency.Seconds())
	}
}

func (s *PrometheusServer) Start() error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", s)
	return http.ListenAndServe(s.addr, mux)
}
