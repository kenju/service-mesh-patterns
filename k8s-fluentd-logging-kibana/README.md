# k8s-fluentd-logging-kibana

Change to the context you use

    # for docker desktop for mac
    kubectl use-context docker-desktop

Check whether RBAC is enabled within your cluster

    kubectl cluster-info dump | grep authorization-mode

Build Docker Image

    docker build -t api-app:test ./api
    docker build -t fluentd-kubernetes-daemonset:v1-debian-elasticsearch ./fluentd

Deploy k8s components

    kubectl apply -f api.yaml
    kubectl apply -f elasticsearch.yaml
    kubectl apply -f fluentd.yaml
    kubectl apply -f kibana.yaml

List pods

    kubectl get pods

List daemonset

    kubectl get ds/fluentd -n kube-system

Describe a pod

    kubectl describe pods api-
    kubectl describe pods fluentd- -n kube-system

Describe a serviceaccount

    kubectl describe serviceaccounts fluentd -n kube-system

Setup port forwarding

    kubectl port-forward api-****** 8080
    kubectl port-forward kibana-*** 5601

Get logs from pods

    kubectl logs api-ff57b4b6f-8tskh
    # or use https://github.com/wercker/stern
    stern api-

Exec bash on pods

    kubectl exec -it fluentd-zw8lx /bin/bash -n kube-system

Delete k8s components

    kubectl delete -f api.yaml

## References
### Fluentd
- https://docs.fluentd.org/v/0.12/articles/kubernetes-fluentd
- https://github.com/fluent/fluentd-kubernetes-daemonset
