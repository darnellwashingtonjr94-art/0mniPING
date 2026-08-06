package l7_app

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func ProbeTLSExpiry(ctx context.Context, target string) telemetry.Result {
	start := time.Now()
	res := telemetry.Result{
		Timestamp: start,
		Target:    target,
		Layer:     "L7",
		Probe:     "tls_expiry",
	}

	dialer := &net.Dialer{Timeout: 3 * time.Second}
	conn, err := tls.DialWithDialer(dialer, "tcp", target, &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}
	defer conn.Close()

	certs := conn.PeerCertificates()
	if len(certs) == 0 {
		res.Status = telemetry.StatusDown
		res.Error = "no peer certificates found"
		return res
	}

	expiry := certs[0].NotAfter
	if time.Now().After(expiry) {
		res.Status = telemetry.StatusDown
		res.Error = "certificate expired"
		return res
	}

	res.Status = telemetry.StatusUp
	res.Latency = time.Since(start)
	return res
}
