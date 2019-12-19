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

## Hands-On

Setup

```bash
# https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/
minikube addons enable ingress
kubectl get pods -n kube-system | grep nginx-ingress-controller

# kubectl krew install ingress-nginx
kubectl krew install ingress-nginx
```

```bash
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
```

## Architecture

### name based virtual hosting

```bash
$ kubectl describe ingress test-ingress
Name:             test-ingress
Namespace:        default
Address:          192.168.64.3
Default backend:  default-http-backend:80 (<none>)
Rules:
  Host         Path  Backends
  ----         ----  --------
  foo.bar.com
                  service1:80 (172.17.0.10:8080,172.17.0.9:8080)
  bar.foo.com
                  service2:80 (172.17.0.10:8080,172.17.0.9:8080)
Annotations:
  kubectl.kubernetes.io/last-applied-configuration:  {"apiVersion":"extensions/v1beta1","kind":"Ingress","metadata":{"annotations":{"nginx.ingress.kubernetes.io/ssl-redirect":"\\\"false\\\""},"name":"test-ingress","namespace":"default"},"spec":{"rules":[{"host":"foo.bar.com","http":{"paths":[{"backend":{"serviceName":"service1","servicePort":80}}]}},{"host":"bar.foo.com","http":{"paths":[{"backend":{"serviceName":"service2","servicePort":80}}]}}]}}

  nginx.ingress.kubernetes.io/ssl-redirect:  \"false\"
Events:
  Type    Reason  Age   From                      Message
  ----    ------  ----  ----                      -------
  Normal  CREATE  41m   nginx-ingress-controller  Ingress default/test-ingress
  Normal  UPDATE  40m   nginx-ingress-controller  Ingress default/test-ingress
```

### path based virtual hosting

```bash
$ kubectl describe ingress test-ingress
Name:             test-ingress
Namespace:        default
Address:
Default backend:  default-http-backend:80 (<none>)
Rules:
  Host         Path  Backends
  ----         ----  --------
  foo.bar.com
               /foo   service1:80 (172.17.0.7:8080,172.17.0.8:8080)
               /bar   service1:80 (172.17.0.7:8080,172.17.0.8:8080)
Annotations:
  kubectl.kubernetes.io/last-applied-configuration:  {"apiVersion":"extensions/v1beta1","kind":"Ingress","metadata":{"annotations":{"nginx.ingress.kubernetes.io/ssl-redirect":"\\\"false\\\""},"name":"test-ingress","namespace":"default"},"spec":{"rules":[{"host":"foo.bar.com","http":{"paths":[{"backend":{"serviceName":"service1","servicePort":80},"path":"/foo"},{"backend":{"serviceName":"service1","servicePort":80},"path":"/bar"}]}}]}}

  nginx.ingress.kubernetes.io/ssl-redirect:  \"false\"
Events:
  Type    Reason  Age   From                      Message
  ----    ------  ----  ----                      -------
  Normal  CREATE  3s    nginx-ingress-controller  Ingress default/test-ingress
```
