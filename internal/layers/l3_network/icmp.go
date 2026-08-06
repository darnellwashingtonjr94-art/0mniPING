package l3_network

import (
	"context"
	"net"
	"time"

	"github.com/darnellwashintonjr94-art/omniping/pkg/telemetry"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func ProbeICMP(ctx context.Context, target string) telemetry.Result {
	start := time.Now()
	res := telemetry.Result{
		Timestamp: start,
		Target:    target,
		Layer:     "L3",
		Probe:     "icmp",
	}

	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}
	defer c.Close()

	peer, err := net.ResolveIPAddr("ip4", target)
	if err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}

	wm := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: os.Getpid() & 0xffff, Seq: 1,
			Data: []byte("omniping-probe"),
		},
	}
	b, err := wm.Marshal(nil)
	if err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}

	if err := c.SetDeadline(time.Now().Add(2 * time.Second)); err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}

	if _, err := c.WriteTo(b, peer); err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}

	rb := make([]byte, 1500)
	_, _, err = c.ReadFrom(rb)
	if err != nil {
		res.Status = telemetry.StatusDown
		res.Error = err.Error()
		return res
	}

	res.Status = telemetry.StatusUp
	res.Latency = time.Since(start)
	return res
}
