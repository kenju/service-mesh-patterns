.PHONY: main
main:
	for i in {1..100}; do curl 'localhost:8081/hello?mirror=1' && sleep 1 ; done

.PHONY: metrics
metrics:
	curl 'localhost:8082/metrics'

.PHONY: mirror
mirror:
	curl 'localhost:8081/hello?mirror=1'
