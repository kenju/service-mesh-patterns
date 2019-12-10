# nginx-proxy-cache

A reverse proxy (NGINX) cache response from upstreams using [ngx_http_proxy_module](https://nginx.org/en/docs/http/ngx_http_proxy_module.html).

## Architecture

```
NGINX (caching response)
`--- backend1 (Go HTTP app)
`--- backend2 (Go HTTP app)
```

## Development

Start containers:

```
make start
```

Then, send HTTP request to the nginx reverse proxy multiple times.

```
for i in {1..10}; do make ping-lb; done
```

You can notice that the following request will not be proxied to backend servers because the first response is cached.

```
backend2_1  | 2019/12/10 13:26:45 backend-service listenining at :8080...
backend1_1  | 2019/12/10 13:26:46 backend-service listenining at :8080...
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:26:48 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
backend1_1  | header: map[Accept:[*/*] User-Agent:[curl/7.54.0] X-Forwarded-Proto:[http] X-Real-Ip:[172.23.0.1]]
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:26:50 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:26:51 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:26:53 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
nginx_1     | 172.23.0.1 - - [10/Dec/2019:13:29:01 +0000] "GET / HTTP/1.1" 200 37 "-" "curl/7.54.0" "-"
```

