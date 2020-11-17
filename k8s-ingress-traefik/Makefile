.PHONY: curl
curl:
	curl "example.com/foo"
	curl "example.com/bar"

.PHONY: apply
apply:
	kubectl apply -k ./overlays/development

.PHONY: delete
delete:
	kubectl delete -k ./overlays/development

.PHONY: start
start:
	minikube start

.PHONY: cleanup
cleanup: delete
	minikube delete
