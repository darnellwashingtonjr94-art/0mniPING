# Omniping

Omniping is a modular, multi-layer network observability and diagnostic suite written in Go. It delivers precise, asynchronous probing across Layer 2 links, Layer 3 networks, Layer 4 transport protocols, Layer 7 applications, DNS resolution, and virtual tunnels, providing real-time console streaming, structured JSON logs, and Prometheus metrics endpoints.

## Features
- **Multi-Layer Probing**: Supports L2 (ARP, OAM), L3 (ICMP, Cloud VPC), L4 (TCP SYN, UDP), L7 (HTTP, gRPC health), DNS, and Tunnels (IPsec DPD, GTP Echo, MPLS LSP).
- **Asynchronous Worker Pool**: High-throughput parallel target execution.
- **Flexible Exporters**: Real-time console output, JSON log streaming, and Prometheus metrics scraping.
- **Configurable Scheduler**: Cron-style interval execution for continuous monitoring.

## Installation & Building
```bash
git clone [https://github.com/darnellwashingtonjr94-art/omniping.git](https://github.com/darnellwashingtonjr94-art/omniping.git)
cd omniping
./scripts/build.sh
