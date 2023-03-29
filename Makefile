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

install-kind:
ifeq ($(shell uname -s), Linux)
	curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.15.0/kind-linux-amd64
	chmod +x ./kind
	sudo mv ./kind /usr/local/bin/kind
else
	brew install kind
endif

kind_check := $(shell command -v kind 2> /dev/null)
setup-kind:
ifndef kind_check
	echo "No kind in $(PATH), installing kind..."
	$(MAKE) install-kind
endif
	echo "kind already insteall.. setting up kind cluster"
	kind create cluster --name $(KIND_CLUSTER)

install-kubectl:
ifeq ($(shell uname -s), Linux)
	curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
	chmod +x ./kubectl
	sudo mv ./kubectl /usr/local/bin/kubectl
else
	brew install kubectl
endif

kubectl_check := $(shell command -v kubectl 2> /dev/null)
setup-kubectl:
ifndef kubectl_check
	$(MAKE) install-kubectl
endif
	echo "kubectl already installed"

setup: setup-kubectl setup-kind

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

.PHONY: port-forward port-forward-all
.port-forward:
	kubectl port-forward pods/$(chain)-genesis-0 $(localrpc):26657 &
	kubectl port-forward pods/$(chain)-genesis-0 $(localrest):1317 &

num_chains = $(shell yq -r ".chains | length" $(VALUES_FILE))
port-forward-all:
	echo "Port forwarding all chains to localhost"
	for i in $(shell seq 0 $(num_chains)); do \
  		$(MAKE) .port-forward \
  			chain=$$(yq -r ".chains[$$i].name" $(VALUES_FILE)) \
  			localrpc=$$(yq -r ".chains[$$i].ports.rpc" $(VALUES_FILE)) \
  			localrest=$$(yq -r ".chains[$$i].ports.rest" $(VALUES_FILE)); \
	done
	echo "Port forwarding explorer to localhost"
	kubectl port-forward service/explorer 8080:8080 &

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

.PHONY: check-forward
check-forward-all:
	while true ; do \
		for port in $$(yq -r "(.chains[] | .ports | .rpc, .rest), .explorer.localPorts.rest" $(VALUES_FILE)); do \
			nc -vz 127.0.0.1 $$port; \
		done; \
		sleep 10; \
	done
