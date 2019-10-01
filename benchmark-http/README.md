# benchmark-http

# Usage

## Loadtesting

Send `GET /` to benchmark-service to start load testing:

```
curl localhost:8001/start | jq .
```

Then see metrics which are supposed to be scraped by Prometheus:

```
curl localhost:8001/metrics
```

output:

```
# HELP benchmark_server_count The total number of completed requests including successful and failed requests
# TYPE benchmark_server_count gauge
benchmark_server_count 200
# HELP benchmark_server_total The total time spent running the test within ghz from start to finish. This is a single measurement from start of the test run to the completion of the final request of the test run.
# TYPE benchmark_server_total gauge
benchmark_server_total 72
```

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

output:

```
               Name                             Command               State                Ports
--------------------------------------------------------------------------------------------------------------
benchmark-http_backend-service_1     /bin/sh -c go run main.go        Up      0.0.0.0:8080->8080/tcp
benchmark-http_benchmark-service_1   /go/benchmark-service            Up      0.0.0.0:8001->8001/tcp, 8080/tcp
benchmark-http_grafana_1             /run.sh                          Up      0.0.0.0:3000->3000/tcp
benchmark-http_prometheus_1          /bin/prometheus --config.f ...   Up      0.0.0.0:9090->9090/tcp
```

# Debug

## prometheus explorer

Open http://localhost:9090

> See docker-compose.yaml

## grafana dashboard

Open http://localhost:3000

> See docker-compose.yaml
