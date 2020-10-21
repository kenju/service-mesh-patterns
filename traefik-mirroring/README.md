# traefik-mirroring

## Usage

```
docker-compose up --build
```

To mirror request, add `mirror=1` request parameter:

```
curl 'localhost:18000/hello?mirror=1'
```
