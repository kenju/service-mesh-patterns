# service-mesh-patterns

The design patterns of container application deployments

## Benchmark

sub directory | description
---|---
benchmark-grpc | load-testing pattern using [bojand/ghz](https://godoc.org/github.com/bojand/ghz) for gRPC load testing
benchmark-http | load-testing pattern using [rakyll/hey](https://github.com/rakyll/hey/) for HTTP load testing

## Envoy Proxy

sub directory | description
---|---
envoy-front-proxy | [front-proxy](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy.html) pattern using [HTTP connection manager filter](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto)
envoy-grpc-gateway | [grpc-bridge](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/grpc_bridge) pattern using [gRPC bridge filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_bridge_filter#config-http-filters-grpc-bridge)
envoy-lua-filters | http-sniffing pattern using [Lua HTTP filter](https://www.envoyproxy.io/docs/envoy/v1.7.0/configuration/http_filters/lua_filter)
envoy-prometheus-monitoring | [distributed-tracing](https://microservices.io/patterns/observability/distributed-tracing.html) pattern using Prometheus & Grafana

## Fluentd

sub directory | description
---|---
fluentd-logging-driver | Basic configuration for [fluentd as a logging driver](https://docs.fluentd.org/container-deployment/docker-compose)
fluentd-logging-kibana | Ingest jsonl format logs from Rails app, parse at fluentd, write messages to Elasticsearch, and query on Kibana

## kubernetes

sub directory | description
---|---
k8s-ingress-nginx | A kubernetes resources using official [ingress-controller](https://kubernetes.io/docs/concepts/services-networking/ingress/) for [NGINX](https://github.com/kubernetes/ingress-nginx)
k8s-fluentd-logging-kibana | Ingest jsonl format logs from Rails app via `/var/log/containers/*.log`, parse at fluentd DaemonSet, write messages to Elasticsearch, and query on Kibana
k8s-fluentd-metrics-prometheus | A kubernetes resources for collecting fluentd metrics via Prometheus
k8s-volume | A minimum example for mounting volumes

## mtail

sub directory | description
---|---
mtail-nginx | The basic architecture for extracting logs from NGINX access logs.

## NGINX

sub directory | description
---|---
nginx-content-distribution | A NGINX template for a scalable architecture for content distribution, which is introduced by ["nginx実践入門"](http://gihyo.jp/magazine/wdpress/plus/978-4-7741-7866-0)
nginx-name-based-virtual-hosting | A NGINX template for the name besed virtual hosting.
nginx-path-based-routing | A NGINX template for the path based routing.
nginx-proxy-cache | A reverse proxy (NGINX) cache response from upstreams using [ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html).
nginx-static-contents | A NGINX template for serving static files.
openresty-simple | A simple container setup based on [openresty/docker-openresty](https://github.com/openresty/docker-openresty).

## Prometheus

sub directory | description
---|---
prometheus-grafana | A basic template for Prometheus for scraping from node_exporter and visualize via the Grafana dashboard.
prometheus-bpf-exporter | A basic architecture for exporting BPF metrics to Prometheus, and visualize on Grafana.
