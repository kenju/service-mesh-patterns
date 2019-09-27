FROM envoyproxy/envoy-dev:latest

RUN apt-get update && apt-get -q install -y \
    curl

COPY ./front-envoy.yaml /etc/front-envoy.yaml

# @doc https://www.envoyproxy.io/docs/envoy/latest/operations/cli
CMD /usr/local/bin/envoy \
    --config-path /etc/front-envoy.yaml \
    --service-cluster front-proxy \
    --service-node front-proxy