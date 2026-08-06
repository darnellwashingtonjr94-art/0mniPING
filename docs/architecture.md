# Omniping Architecture

Omniping is structured around a decoupled multi-layer monitoring framework designed for high-concurrency network observability.

## Component Layout
- **Engine**: Coordinates asynchronous worker pools, rate limiters, cron-style schedulers, and probe dispatchers.
- **Layers**: Implements isolated protocol probes across OSI layers 2 through 7, including specialized tunnel and DNS diagnostics.
- **Exporters**: Pluggable telemetry output interfaces supporting Console (TTY), JSON log rotation, Prometheus endpoints, and Syslog.
- **Telemetry**: Common data structures, metric aggregators, and result buffers.
