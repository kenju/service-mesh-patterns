# nginx-scalable-lb-for-content-distribution

a scalable architecture for content distribution, which is introduced by ["nginx実践入門"](http://gihyo.jp/magazine/wdpress/plus/978-4-7741-7866-0) by [@cubicdaiya](https://github.com/cubicdaiya).

## Architecture

![](https://i.gyazo.com/e808a0e196f8797eefc5398d65374bbb.png)
> Credit: ["nginx実践入門"](http://gihyo.jp/magazine/wdpress/plus/978-4-7741-7866-0) by [@cubicdaiya](https://github.com/cubicdaiya) Chapter 7, Section 6


## Development

Start containers:

```
make start
```

Then, send HTTP request to the nginx reverse proxy multiple times.

```
make load-test
```

You can find that the origin server is accessed only for the first time. Following request are load balanced & hit from cache servers.

```
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:57:56 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:57:56 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
origin_1  | 172.24.0.4 - - [10/Dec/2019:13:57:56 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:58:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:58:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:58:02 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:58:02 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:58:03 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:58:03 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:00 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
lb_1      | 172.24.0.1 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
cache1_1  | 172.24.0.5 - - [10/Dec/2019:13:59:01 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.54.0" "-"
```
