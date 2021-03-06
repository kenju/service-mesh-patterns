# prometheus-grafana

The basic architecture for exporting BPF metrics to Prometheus, and visualize on Grafana.

# Usage

## Development

Run containers:

```
docker-compose pull
docker-compose build
docker-compose up
```

Check the running containers:

```
docker-compose ps
```

## Containers

### bpf-exporter

The BPF scripts copyrights for https://github.com/cloudflare/ebpf_exporter

### prometheus explorer

Open http://localhost:9090

### grafana dashboard

Open http://localhost:3000

Default username/password is `admin`.

> On the login page, type admin for the username and password.
> https://grafana.com/docs/grafana/latest/getting-started/getting-started/#log-in-for-the-first-time
