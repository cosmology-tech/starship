# Set helm args based on only file to show
HELM_NAME = starship
HELM_CHART = devnet
ifneq ($(FILE),)
HELM_ARGS += --show-only $(FILE)
endif

KEYS_CONFIG = charts/$(HELM_CHART)/configs/keys.json
VALUES_FILE = charts/$(HELM_CHART)/values.yaml

###############################################################################
###                              Proto commands                             ###
###############################################################################
.PHONY: proto
proto:
	(cd proto/ && make build-proto)

###############################################################################
###                              Helm commands                              ###
###############################################################################

all: delete install

debug:
	helm template --dry-run --debug --generate-name ./charts/$(HELM_CHART) -f $(VALUES_FILE) $(HELM_ARGS)

install:
	helm install --replace --debug $(HELM_NAME) ./charts/$(HELM_CHART) -f $(VALUES_FILE) $(HELM_ARGS)

upgrade:
	helm upgrade --debug $(HELM_NAME) ./charts/$(HELM_CHART) -f $(VALUES_FILE) $(HELM_ARGS)

test:
	helm test --debug $(HELM_NAME)

delete:
	-helm delete --debug $(HELM_NAME)

###############################################################################
###                          Local Kind Setup                               ###
###############################################################################
KIND_CLUSTER=starship

.PHONY: check
check:
	bash $(CURDIR)/scripts/dev-setup.sh

.PHONY: setup
setup: setup-kind

.PHONY: clean
clean: clean-kind

.PHONY: setup-kind
setup-kind:
	kind create cluster --name $(KIND_CLUSTER)

.PHONY: clean-kind
clean-kind:
	kind delete cluster --name $(KIND_CLUSTER)

###############################################################################
###                              Keys config                                ###
###############################################################################

# Helper scripts to play with the keys
# Requires a chain binary to be installed locally
# one of validators, keys, relayers
TYPE = relayers
N = 1

number = $(shell jq '.$(TYPE) | length + 1' $(KEYS_CONFIG))
name = $(shell echo $(shell jq '.$(TYPE)[0].name' $(KEYS_CONFIG) | sed 's/[0-9]//'))$(number)
mnemonic = $(shell gaiad keys add rly --keyring-backend test --dry-run --output json 2>&1 | jq -r ".mnemonic")
add-mnemonic:
	@jq '.$(TYPE) += [{"name": "$(name)", "type": "local", "mnemonic": "$(mnemonic)"}]' $(KEYS_CONFIG) >> $(KEYS_CONFIG).temp
	@mv $(KEYS_CONFIG).temp $(KEYS_CONFIG)
	jq -r '.$(TYPE) | last' $(KEYS_CONFIG)

add-n-mnemonic:
	i=1; while [ "$$i" -le $(N) ]; do \
		$(MAKE) add-mnemonic; i=$$((i + 1));\
	done

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHOY: port-forward-all
port-forward-all:
	$(CURDIR)/scripts/port-forward.sh --config=$(VALUES_FILE)

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"
