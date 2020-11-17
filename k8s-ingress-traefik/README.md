# k8s-ingress-traefik

https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/

## Documentation

- https://kubernetes.io/docs/concepts/services-networking/ingress/
- https://doc.traefik.io/traefik/routing/providers/kubernetes-ingress/
- https://doc.traefik.io/traefik/v2.3/user-guides/crd-acme/#cluster-resources

## Deploy

```
kubectl apply -k ./overlays/development
kubectl delete -k ./overlays/development
```
