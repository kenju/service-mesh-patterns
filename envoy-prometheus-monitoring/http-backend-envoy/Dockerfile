FROM envoyproxy/envoy-dev:latest

RUN apt-get update && apt-get -q install -y \
    curl

COPY ./backend-envoy.yaml /etc/backend-envoy.yaml

CMD /usr/local/bin/envoy \
    --config-path /etc/backend-envoy.yaml \
    --service-cluster backend-envoy \
    --service-node backend-envoy