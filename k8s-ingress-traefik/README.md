# k8s-ingress-traefik

https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/

## Documentation

- https://kubernetes.io/docs/concepts/services-networking/ingress/
- https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/
- https://doc.traefik.io/traefik/v2.3/user-guides/crd-acme/#cluster-resources

## Start

### Minikube

- https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/

At first, start the minikube process and enable the ingress [addon](https://minikube.sigs.k8s.io/docs/handbook/deploying/).

```
minikube start --vm=true
minikube addons enable ingress
```

Assuming that `Ingress.spec.rules[0].host` is set to `example.com` (see `ingress.yaml`), update the static host lookup table with the following:

```
echo "$(minikube ip) example.com" >> /etc/hosts
```

Now the following request would return from the [whoami](https://github.com/traefik/whoami) services.

```
curl "http://example.com/foo"
curl "http://example.com/bar"
```

## Deploy

```
kubectl apply -k ./overlays/development
kubectl delete -k ./overlays/development
```
