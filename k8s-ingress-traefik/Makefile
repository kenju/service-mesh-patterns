.PHONY: apply
apply:
	kubectl apply -k ./overlays/development

.PHONY: delete
delete:
	kubectl delete -k ./overlays/development

.PHONY: start
start:
	minikube start --vm=true
	minikube addons enable ingress

.PHONY: cleanup
cleanup: delete
	minikube delete
