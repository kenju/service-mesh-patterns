FROM envoyproxy/envoy-dev:latest

RUN apt-get update && apt-get -q install -y \
    curl

COPY ./front-envoy.yaml /etc/front-envoy.yaml

CMD /usr/local/bin/envoy \
    -c /etc/front-envoy.yaml \
    --service-cluster front-proxy \
    --service-node backend