package engine

import (
	"context"

	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/dns"
	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l2_link"
	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l3_network"
	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l4_transport"
	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/l7_app"
	"github.com/darnellwashingtonjr94-art/omniping/internal/layers/tunnel"
	"github.com/darnellwashingtonjr94-art/omniping/pkg/telemetry"
)

func DispatchProbe(ctx context.Context, layer, probeType, target string) telemetry.Result {
	switch layer {
	case "L2":
		if probeType == "arping" {
			return l2_link.ProbeARP(ctx, target)
		}
		return l2_link.ProbeOAMEth(ctx, target)
	case "L3":
		if probeType == "cloud_vpc" {
			return l3_network.ProbeCloudVPC(ctx, target)
		}
		return l3_network.ProbeICMP(ctx, target)
	case "L4":
		if probeType == "udp_probe" {
			return l4_transport.ProbeUDP(ctx, target)
		}
		return l4_transport.ProbeTCPSYN(ctx, target)
	case "L7":
		if probeType == "grpc_health" {
			return l7_app.ProbeGRPCHealth(ctx, target)
		}
		return l7_app.ProbeHTTPHealth(ctx, target)
	case "DNS":
		return dns.ProbeDNSResolver(ctx, target)
	case "Tunnel":
		if probeType == "gtp_echo" {
			return tunnel.ProbeGTPEcho(ctx, target)
		}
		if probeType == "mpls_lsp" {
			return tunnel.ProbeMPLSLSP(ctx, target)
		}
		return tunnel.ProbeIPSecDPD(ctx, target)
	default:
		return l3_network.ProbeICMP(ctx, target)
	}
}
