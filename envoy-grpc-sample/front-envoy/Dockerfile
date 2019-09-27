FROM envoyproxy/envoy:latest

COPY front-envoy.yaml /etc/front-envoy.yaml

RUN apt-get update && apt-get -q install -y \
    curl

CMD /usr/local/bin/envoy -c /etc/front-envoy.yaml --service-cluster front-proxy