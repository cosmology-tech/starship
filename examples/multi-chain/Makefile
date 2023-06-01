NAME = starship-multichain
FILE = config.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.1.31

###############################################################################
###                              All commands                               ###
###############################################################################

.PHONY: setup
setup: check setup-helm

.PHONY: start
start: install port-forward

.PHONY: stop
stop: stop-forward delete

.PHONY: clean
clean: stop clean-kind

###############################################################################
###                              Helm Charts                                ###
###############################################################################

setup-helm:
	helm repo add $(HELM_REPO) https://cosmology-tech.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

install:
	helm install -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION) --wait --timeout 20m

upgrade:
	helm upgrade --debug $(NAME) $(HELM_REPO)/$(HELM_CHART) -f $(FILE) --version $(HELM_VERSION)

debug:
	helm install --dry-run --debug -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

delete:
	-helm delete $(NAME)

###############################################################################
###                                 Port forward                            ###
###############################################################################

.PHONY: port-forward
port-forward:
	bash $(CURDIR)/../../scripts/port-forward.sh --config=$(FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

###############################################################################
###                          Dependency check                               ###
###############################################################################

.PHONY: check
check:
	bash $(CURDIR)/../../scripts/dev-setup.sh

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
