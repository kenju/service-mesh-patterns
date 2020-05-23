#!/usr/bin/env sh

sudo docker run \
  -it \
  --rm \
  --privileged \
  --volume /lib/modules:/lib/modules:ro \
  --volume /usr/src:/usr/src:ro \
  --volume /etc/localtime:/etc/localtime:ro \
  -p 9435:9435 \
  kenju/bpf-exporter \
  /bin/bash
