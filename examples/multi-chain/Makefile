NAME = multi-chain
FILE = configs/starship.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.2.1

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

# run command with `$ make install FILE=configs/osmo-juno-cosmos.yaml`
.PHONY: install
install:
	bash $(CURDIR)/scripts/install.sh --config $(FILE) --name $(NAME) --version $(HELM_VERSION)

.PHONY: install-base
install-base:
	bash $(CURDIR)/scripts/install.sh --config $(FILE) --name $(NAME) --version $(HELM_VERSION) --chart $(CURDIR)/../../charts/devnet

delete:
	-helm delete $(NAME)

###############################################################################
###                                 Port forward                            ###
###############################################################################

# run command with `$ make port-forward FILE=configs/osmo-juno-cosmos.yaml`
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
