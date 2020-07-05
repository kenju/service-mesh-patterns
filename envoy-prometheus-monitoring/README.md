# envoy-prometheus-monitoring

The idea of this repository comes from https://github.com/dnivra26/envoy_monitoring

## Architecture

![](./envoy-prometheus-monitoring.jpg)

## Usage

### docker-compose

Run containers:

    docker-compose up --build

Send request to front-envoy, and see the response:

    curl localhost:9000

### front-envoy admin

Open http://localhost:8001

### http-backend-envoy admin

Open http://localhost:8091

### prometheus explorer

Open http://localhost:9090

### grafana dashboard

Open http://localhost:3000

### statsd_exporter

Open http://localhost:9102
