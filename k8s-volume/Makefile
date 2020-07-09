.PHONY: build
build:
	docker build -t fluentd-kubernetes-daemonset:k8s-volume ./fluentd

.PHONY: apply
apply:
	kubectl apply -k ./overlays/development

.PHONY: delete
delete:
	kubectl delete -k ./overlays/development

.PHONY: curl
curl:
	curl localhost:8080/health

.PHONY: ping
ping:
	for i in {0..1000}; do curl localhost:8080/health; sleep 3; done
