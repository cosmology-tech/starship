NAME = cosmos-injective-cluster
FILE = configs/local.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.1.39

###############################################################################
###                              All commands                               ###
###############################################################################

.PHONY: setup
setup: setup-deps setup-helm setup-kind

.PHONY: stop
stop: stop-forward delete

.PHONY: clean
clean: stop clean-kind

###############################################################################
###                          Dependency check                               ###
###############################################################################

.PHONY: check
setup-deps:
	bash $(CURDIR)/scripts/dev-setup.sh

###############################################################################
###                              Helm Charts                                ###
###############################################################################

setup-helm:
	helm repo add $(HELM_REPO) https://cosmology-tech.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

install:
	helm install -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

install-dev:
	helm install -f $(FILE) $(NAME) ../../charts/$(HELM_CHART)

delete:
	-helm delete $(NAME)

###############################################################################
###                               Custom Docker                             ###
###############################################################################
DOCKER_REPO := public.ecr.aws

.PHONY: docker-build
docker-build:
	docker buildx build --platform linux/amd64 -t $(DOCKER_REPO)/a6c1y1k3/injective:v1.11.5 . --push

.PHONY: docker-run
docker-run:
	docker run --rm -it --entrypoint /bin/bash $(DOCKER_REPO)/a6c1y1k3/injective:v1.11.5

###############################################################################
###                                 Port forward                            ###
###############################################################################

.PHONY: port-forward
port-forward:
	bash $(CURDIR)/scripts/port-forward.sh --config=$(FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

###############################################################################
###                          Local Kind Setup                               ###
###############################################################################
KIND_CLUSTER=starship

.PHONY: setup-kind
setup-kind:
	kind create cluster --name $(KIND_CLUSTER)

.PHONY: clean-kind
clean-kind:
	kind delete cluster --name $(KIND_CLUSTER)
