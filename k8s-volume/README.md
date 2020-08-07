# k8s-volume

Change to the context you use

    kubectl use-context docker-desktop # for docker desktop for mac

Check whether RBAC is enabled within your cluster

    kubectl cluster-info dump | grep authorization-mode

Build Docker Image

    make build

Deploy k8s components

    make apply

Describe a pod

    kubectl describe pods api-
    kubectl describe pods fluentd- -n kube-system

Describe a serviceaccount

    kubectl describe serviceaccounts fluentd -n kube-system

Setup port forwarding

    kubectl port-forward api-****** 8080

Make cURL request

    make curl

Get logs from pods

    kubectl logs api-ff57b4b6f-8tskh
    # or use https://github.com/wercker/stern
    stern api-

Exec bash on pods

    kubectl exec -it fluentd-zw8lx /bin/bash -n kube-system

Delete k8s components

    make delete
