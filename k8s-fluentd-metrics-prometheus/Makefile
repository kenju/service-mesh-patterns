.PHONY: build
build:
	docker build -t fluentd-kubernetes-daemonset:k8s-fluentd-metrics-prometheus ./fluentd

.PHONY: apply
apply:
	kubectl apply -k ./overlays/development

.PHONY: delete
delete:
	kubectl delete -k ./overlays/development
