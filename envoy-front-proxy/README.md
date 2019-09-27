# envoy-front-proxy-sample

https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/front_proxy.html

Start containers:

```
docker-compose pull
docker-compose build
docker-compose --compatibility up
```

Test connection from other process:

```
docker-compose ps
curl -v localhost:8000/service/1
curl -v localhost:8000/service/2
```

Scale up and chech that the request from front to the backend services will be round-robin:

```
docker-compose scale service1=3
curl -v localhost:8000/service/1
curl -v localhost:8000/service/1
curl -v localhost:8000/service/1
```

Enter front container and see the request to the backend serrvices:

```
docker-compose exec front-envoy /bin/bash
> curl localhost:80/service/1
> curl localhost:80/service/1
> curl localhost:80/service/2
> curl localhost:80/service/2
```

See admin metrics for front envoy:

```
docker-compose exec front-envoy /bin/bash
> curl localhost:8001/server_info
```

Or visit the admin page from a browser:

```
open http://localhost:8001/
```
