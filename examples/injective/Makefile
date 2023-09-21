NAME = injective
FILE = configs/local.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.1.46

###############################################################################
###                              Helm Charts                                ###
###############################################################################

.PHONY: start
start: install port-forward

.PHONY: stop
stop: stop-forward delete

###############################################################################
###                              Helm Charts                                ###
###############################################################################

install:
	bash $(CURDIR)/scripts/install.sh --config $(FILE) --name $(NAME) --version $(HELM_VERSION)

debug:
	bash $(CURDIR)/scripts/install.sh --config $(FILE) --name $(NAME) --dry-run --version $(HELM_VERSION)

delete:
	-helm delete $(NAME)

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
