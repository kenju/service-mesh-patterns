# containers-k8s-ingress-nginx

## Resources

Documentation

- https://github.com/kubernetes/ingress-nginx
- https://kubernetes.github.io/ingress-nginx/deploy/

Guide

- https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/
- https://www.youtube.com/watch?v=AXZr2OC8Unc (from 11:28~)
- https://medium.com/@awkwardferny/getting-started-with-kubernetes-ingress-nginx-on-minikube-d75e58f52b6c

Image

- https://github.com/kubernetes/kubernetes/tree/b1766b707a0f61e4a2640359c4482dcb2c689df5/test/images/echoserver
- https://console.cloud.google.com/gcr/images/kubernetes-e2e-test-images/GLOBAL/echoserver?gcrImageListsize=30

## Usage

    # kubectl krew install ingress-nginx
    kubectl krew install ingress-nginx
    kubectl krew list

    # create deployment, pods, services and ingress
    kubectl apply -k overlays/name-baed
    kubectl apply -k overlays/path-baed

    # port-forward to services
    kubectl port-forward test-*** 8081:8080
    kubectl port-forward test-*** 8082:8080

    # curl to services
    curl localhost:8081/foo
    curl localhost:8082/bar

    # describe
    kubectl describe ingress test-ingress

    # delete
    kubectl delete -k overlays/name-baed
    kubectl delete -k overlays/path-baed

### Minikube

#### Setup

    # https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/
    minikube addons enable ingress
    kubectl get pods -n kube-system | grep nginx-ingress-controller

    # kubectl krew install ingress-nginx
    kubectl krew install ingress-nginx

#### Deployment

    # check minikube is running
    kubectl get nodes
    kubectl describe node minikube

    # create deployment, pods, services and ingress
    kubectl apply -f app.yaml

    # check each resources are created
    kubectl get deploy
    kubectl get pods
    kubectl get service
    kubectl get ingress

#### Operation

    # launch minikube dashboard
    minikube dashboard

    # accessing the application
    minikube ip

    # cURL base.yaml
    curl `minikube ip`/test

    # cURL name-based-virtual-hosting.yaml
    curl -H "Host: foo.bar.com" `minikube ip`
    curl -H "Host: bar.foo.com" `minikube ip`

    # cURL path-based-routing.yaml
    curl -H "Host: foo.bar.com" `minikube ip`/foo
    curl -H "Host: foo.bar.com" `minikube ip`/bar

    # describe pods
    kubectl get pods -n kube-system \
    | grep nginx-ingress-controller \
    | awk '{print $1}' \
    | kubectl describe pods -n kube-system

    # describe ingress
    kubectl describe ingress test-ingress

    # view the logs
    kubectl get pods -n kube-system \
    | grep nginx-ingress-controller \
    | awk '{print $1}' \
    | xargs kubectl logs -n kube-system

    # view the nginx conf
    kubectl get pods -n kube-system \
    | grep nginx-ingress-controller \
    | awk '{print $1}' \
    | xargs -J % kubectl exec -it -n kube-system % cat /etc/nginx/nginx.conf

    # get a JSON array of the backends for ingress-nginx
    # https://kubernetes.github.io/ingress-nginx/kubectl-plugin/#backends
    kubectl ingress-nginx backends -n kube-system

    # dump the generated nginx.conf
    # https://kubernetes.github.io/ingress-nginx/kubectl-plugin/#conf
    kubectl ingress-nginx conf -n kube-system --host foo.bar.com

    # list ingress definitions
    # https://kubernetes.github.io/ingress-nginx/kubectl-plugin/#ingresses
    kubectl ingress-nginx ingresses --all-namespaces

    # save the nginx conf to tmp
    kubectl exec \
    -it \
    -n kube-system \
    `kubectl get pods -n kube-system | grep nginx-ingress-controller | awk '{print $1}'` \
    cat /etc/nginx/nginx.conf \
    > tmp/tmp.conf

    # exec bash
    kubectl exec \
    -it \
    -n kube-system \
    `kubectl get pods -n kube-system | grep nginx-ingress-controller | awk '{print $1}'` \
    /bin/bash

    # delete
    kubectl delete -f app.yaml
