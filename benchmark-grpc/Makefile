.PHONY: show-diagram
show-diagram:
	dot architecture.dot -Tjpg -Gsize=9,15\! -Gdpi=60 | imgcat

.PHONY: ouput-diagram
output-diagram:
	dot architecture.dot -Tjpg -Gsize=9,15\! -Gdpi=60 > grpc-benchmark.jpg

.PHONY: start
start:
	curl localhost:8001/start | jq .

.PHONY: metrics
metrics:
	curl localhost:8001/metrics
