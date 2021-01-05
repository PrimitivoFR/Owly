# Kubernetes

`minikube start`
`minikube mount $PWD:/owly`

## Secrets

### Keycloak

`kubectl create secret generic keycloak-creds --from-env-file=keycloak/.keycloak-env`