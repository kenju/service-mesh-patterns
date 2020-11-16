# k8s-ingress-traefik

https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/

## Documentation

- https://kubernetes.io/docs/concepts/services-networking/ingress/
- https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/

## Install

- https://doc.traefik.io/traefik/getting-started/install-traefik/#use-the-helm-chart

```
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
helm install traefik traefik/traefik
```

## Deploy

```
kubectl apply -k ./overlays/development
kubectl delete -k ./overlays/development
```
