# service-mesh-patterns

## Benchmark

sub directory | description
---|---
benchmark-grpc | load-testing pattern using [bojand/ghz](https://godoc.org/github.com/bojand/ghz) for gRPC load testing
benchmark-http | load-testing pattern using [rakyll/hey](https://github.com/rakyll/hey/) for HTTP load testing

## envoy proxy

sub directory | description
---|---
envoy-front-proxy | [front-proxy](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy.html) pattern using [HTTP connection manager filter](https://www.envoyproxy.io/docs/envoy/latest/api-v2/config/filter/network/http_connection_manager/v2/http_connection_manager.proto)
envoy-grpc-gateway | [grpc-bridge](https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/grpc_bridge) pattern using [gRPC bridge filter](https://www.envoyproxy.io/docs/envoy/latest/configuration/http/http_filters/grpc_http1_bridge_filter#config-http-filters-grpc-bridge)
envoy-lua-filters | http-sniffing pattern using [Lua HTTP filter](https://www.envoyproxy.io/docs/envoy/v1.7.0/configuration/http_filters/lua_filter)
envoy-prometheus-monitoring | [distributed-tracing](https://microservices.io/patterns/observability/distributed-tracing.html) pattern using Prometheus & Grafana

## NGINX

sub directory | description
---|---
nginx-proxy-cache | A reverse proxy (NGINX) cache response from upstreams using [ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html).
