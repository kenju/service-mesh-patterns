# benchmark-grpc

# Architecture

![](./grpc-benchmark.jpg)

## Usage

Send `GET /` to benchmark-service to start load testing:

    curl localhost:8001/start | jq .

Then see metrics which are supposed to be scraped by Prometheus:

    curl localhost:8001/metrics

output:

    # HELP benchmark_server_count The total number of completed requests including successful and failed requests
    # TYPE benchmark_server_count gauge
    benchmark_server_count 200
    # HELP benchmark_server_total The total time spent running the test within ghz from start to finish. This is a single measurement from start of the test run to the completion of the final request of the test run.
    # TYPE benchmark_server_total gauge
    benchmark_server_total 72

## Development

Run containers:

    docker-compose up --build

Check the running containers:

    docker-compose ps

output:

                Name                       Command         State           Ports
    ---------------------------------------------------------------------------------------
    benchmark-grpc_backend-service_1   /go/backend-service   Up      0.0.0.0:8080->8080/tcp

Send gRPC request to gateway-service:

    (cd backend-service && make run-client)

output:

    2019/09/28 14:38:18 backend.Hello() message=success:<status_code:"0" >

containers' log:

    backend-service_1  | time="2019-09-28T05:38:18Z" level=info msg="Hello()" func="main.(*backendServer).Hello" file="/app/main.go:88" request=

# Debug

## prometheus explorer

Open http://localhost:9090

## grafana dashboard

Open http://localhost:3000
