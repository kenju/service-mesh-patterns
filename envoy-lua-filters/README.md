# envoy-lua-filters-sample

## Motivation

Logging HTTP POST body using [Lua's HTTP Filter](https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_filters/lua_filter#config-http-filters-lua-stream-handle-api).

PoC for this issue: https://github.com/envoyproxy/envoy/issues/4724

## Implementation

- `envoy`
  - cURL request to this envoy container, then it will proxy request to the `app` container
  - A proxy where Lua's HTTP Filter comes in, which print out HTTP body to stdout
- `app`
  - simple http-echo server
- `log-service`
  - build from sinatra, but this can be anything (e.g. fluentd, RDBMS, NoSQL, File systems, ...)
  - `envoy` container will hook HTTP request, then does HTTP request using [`httpCall()`](https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_filters/lua_filter#httpcall)

### NOTE

This repository was initially based on [envoy's official Lua HTTP Filter sample](https://github.com/envoyproxy/envoy/tree/master/examples/lua).

## Guidelines

Run docker containers as follows:

```
make run
```

Send cURL request to local app container:

```
make curl
```

Look for the log output as follows:

```
output:
envoy_1        | [2019-09-20 13:34:38.503][19][debug][lua] [source/extensions/filters/http/lua/lua_filter.cc:589] script log: {
envoy_1        |   "path": "/",
envoy_1        |   "headers": {
envoy_1        |     "host": "localhost:8000",
envoy_1        |     "user-agent": "curl/7.54.0",
envoy_1        |     "accept": "*/*",
envoy_1        |     "x-forwarded-proto": "http",
envoy_1        |     "x-request-id": "7bbcd056-bb7e-4c24-8f32-9abb5f7c45d3",
envoy_1        |     "x-envoy-expected-rq-timeout-ms": "15000",
envoy_1        |     "content-length": "0"
envoy_1        |   },
envoy_1        |   "method": "GET",
envoy_1        |   "body": "",
envoy_1        |   "fresh": false,
envoy_1        |   "hostname": "localhost",
envoy_1        |   "ip": "::ffff:172.19.0.4",
envoy_1        |   "ips": [],
envoy_1        |   "protocol": "http",
envoy_1        |   "query": {},
envoy_1        |   "subdomains": [],
envoy_1        |   "xhr": false,
envoy_1        |   "os": {
envoy_1        |     "hostname": "129329f6b030"
envoy_1        |   }
envoy_1        | }
```

This request is proxied to log-service using [httpCall()](https://www.envoyproxy.io/docs/envoy/v1.9.0/configuration/http_filters/lua_filter#httpcall) as well:

```
log-service_1  | I, [2019-09-20T13:34:38.508823 #1]  INFO -- : "{\n  \"path\": \"/\",\n  \"headers\": {\n    \"host\": \"localhost:8000\",\n    \"user-agent\": \"curl/7.54.0\",\n    \"accept\": \"*/*\",\n    \"x-forwarded-proto\": \"http\",\n    \"x-request-id\": \"7bbcd056-bb7e-4c24-8f32-9abb5f7c45d3\",\n    \"x-envoy-expected-rq-timeout-ms\": \"15000\",\n    \"content-length\": \"0\"\n  },\n  \"method\": \"GET\",\n  \"body\": \"\",\n  \"fresh\": false,\n  \"hostname\": \"localhost\",\n  \"ip\": \"::ffff:172.19.0.4\",\n  \"ips\": [],\n  \"protocol\": \"http\",\n  \"query\": {},\n  \"subdomains\": [],\n  \"xhr\": false,\n  \"os\": {\n    \"hostname\": \"129329f6b030\"\n  }\n}"
log-service_1  | 172.19.0.4 - - [20/Sep/2019:13:34:38 +0000] "POST /log HTTP/1.1" 200 - 0.0052
```
