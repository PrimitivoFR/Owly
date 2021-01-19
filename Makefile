all: minikube-start minikube-mount k-secrets

minikube-start:
	minikube stop && minikube start --cpus 4 --memory 8192

minikube-mount:
	minikube mount $PWD:/owly

k-secrets:
	kubectl create secret generic keycloak-creds --from-env-file=keycloak/.keycloak-env