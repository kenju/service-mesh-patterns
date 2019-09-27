FROM envoyproxy/envoy-dev:latest

COPY ./lib/mylibrary.lua /lib/mylibrary.lua
COPY ./envoy.yaml /etc/envoy.yaml

CMD /usr/local/bin/envoy -c /etc/envoy.yaml -l debug --service-cluster proxy
