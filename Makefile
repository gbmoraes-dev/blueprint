include .env
export

.PHONY: up down cluster-up cluster-down install-argocd setup-gh-token wait-argocd argocd-password app-of-apps cluster

up:
	@docker compose up --build -d

down:
	@docker compose down -v --remove-orphans

cluster-up:
	@k3d cluster create --config infra/k8s/k3d/k3d.yml

install-argocd:
	@kubectl create namespace argocd
	@kubectl apply -n argocd --server-side -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml

setup-gh-token:
	@kubectl create secret generic ghcr-credentials --from-literal=username=$(GH_USERNAME) --from-literal=password=$(GH_TOKEN) -n argocd
	@kubectl create secret generic git-credentials --from-literal=username=$(GH_USERNAME) --from-literal=password=$(GH_TOKEN) -n argocd

wait-argocd:
	@echo "Aguardando ArgoCD ficar pronto..."
	@kubectl wait --for=condition=available deployment/argocd-server -n argocd --timeout=120s

argocd-password:
	@kubectl -n argocd get secret argocd-initial-admin-secret -o jsonpath="{.data.password}" | base64 -d && echo

app-of-apps:
	@kubectl apply --server-side -f infra/k8s/argocd/app-of-apps.yml

cluster: cluster-up install-argocd setup-gh-token wait-argocd argocd-password app-of-apps

cluster-down:
	@k3d cluster delete blueprint
