# fluentd-logging-via-file

https://docs.fluentd.org/container-deployment/docker-compose

## Development

At first, start up containers:

```
docker-compose up --build
```

Send some HTTP request:

```
for i in {0..100}; do curl -XGET localhost:8080/health; sleep 1; done
```

Check the log from bind-mounted files:

```
docker-compose run fluentd tail -f /fluentd/log/data.log
```

You can directly send logs to the fluentd container for debugging:

```
docker run --log-driver=fluentd --log-opt tag="docker.{{.ID}}" --log-opt fluentd-address=0.0.0.0:24224 python:alpine echo Hello
```

Go to `http://localhost:5601`, create index pattern with `fluentd-*` pattern, and try out indices from query dathboard.
