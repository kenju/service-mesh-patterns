.PHONY: ping-lb
ping-lb:
	curl --verbose --header 'Host: load-balancer' localhost:80

.PHONY: ping-stats
ping-stats:
	curl --verbose localhost:80/stats

.PHONY: start
start:
	docker-compose up --build

