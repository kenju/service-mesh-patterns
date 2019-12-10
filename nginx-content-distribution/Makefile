.PHONY: ping-lb
ping-lb:
	curl --verbose --header 'Host: load-balancer' localhost:80

.PHONY: load-test
load-test:
	for i in {1..10}; do curl -s -o /dev/null -v --header 'Host: load-balancer' localhost:80; done

.PHONY: ping-stats
ping-stats:
	curl --verbose localhost:80/stats

.PHONY: start
start:
	docker-compose up --build

