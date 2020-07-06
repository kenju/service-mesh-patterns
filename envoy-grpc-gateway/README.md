# envoy-grpc-sample

Original: https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/grpc_bridge

## Usage

Run containers:

    docker-compose up --build

Check the running containers:

    docker-compose ps

output:

                Name                              Command               State                      Ports
    --------------------------------------------------------------------------------------------------------------------------
    envoy-grpc-sample_backend-service_1   /go/backend-service              Up      0.0.0.0:8080->8080/tcp
    envoy-grpc-sample_gateway-service_1   /go/gateway-service              Up      0.0.0.0:3000->3000/tcp
    front-proxy                           /docker-entrypoint.sh /bin ...   Up      10000/tcp, 8000/tcp, 0.0.0.0:8001->8001/tcp

Send HTTP request to gateway-service:

    curl localhost:3000/v1/hello

Scale out

    docker-compose up --scale backend-service=3

Browse the front-envoy's admin page:

    open localhost:8001
