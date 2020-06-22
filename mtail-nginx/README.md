# mtail-nginx

The basic architecture for extracting logs from NGINX access logs.

## Setup

Build google/mtail locally at first.

    git clone git@github.com:google/mtail.git
    cd mtail
    docker build -t mtail .

http://google.github.io/mtail/Building.html#no-go

## Development

Run containers:

    docker-compose pull
    docker-compose build
    docker-compose up

Check the running containers:

    docker-compose ps

Send GET request to the NGINX container:

    for i in {0..100}; do curl localhost:8888; sleep 3; done

Check extracted metrics

    curl --silent localhost:3903/metrics | grep 'prog="nginx_access_log.mtail"'

## GUI Console

### mtail

Open http://localhost:3903

### prometheus explorer

Open http://localhost:9090

### grafana dashboard

Open http://localhost:3000

Default username/password is `admin`.

> On the login page, type admin for the username and password.
> https://grafana.com/docs/grafana/latest/getting-started/getting-started/#log-in-for-the-first-time

The `node-exporter-dashboard.json` is downloaded from https://grafana.com/grafana/dashboards/1860.
