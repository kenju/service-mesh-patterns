.PHONY: show-diagram
show-diagram:
	dot architecture.dot -Tjpg -Gsize=9,15\! -Gdpi=60 | imgcat

.PHONY: ouput-diagram
output-diagram:
	dot architecture.dot -Tjpg -Gsize=9,15\! -Gdpi=60 > envoy-prometheus-monitoring.jpg

.PHONY: curl
curl:
	curl localhost:8000
