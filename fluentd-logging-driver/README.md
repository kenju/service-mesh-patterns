# fluentd-logging-driver

## Development

At first, start up containers:

```
docker-compose up --build
```

Send some HTTP request:

```
curl -XGET localhost:4567/
curl -XGET localhost:4567/error
curl -XGET localhost:4567/foo
```

Check the log from bind-mounted files:

```
tail -f ./log/fluentd/*.log
```

You can directly send logs to the fluentd container for debugging:

```
docker run --log-driver=fluentd --log-opt tag="docker.{{.ID}}" --log-opt fluentd-address=0.0.0.0:24224 python:alpine echo Hello
```
