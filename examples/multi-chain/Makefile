NAME = starship-multichain
FILE = config.yaml

HELM_REPO = starship
HELM_CHART = devnet
HELM_VERSION = v0.1.15

###############################################################################
###                              Helm Charts                                ###
###############################################################################

setup:
	helm repo add $(HELM_REPO) https://anmol1696.github.io/starship/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

install:
	helm install -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

upgrade:
	helm upgrade --debug $(NAME) $(HELM_REPO)/$(HELM_CHART) -f $(FILE) --version $(HELM_VERSION)

debug:
	helm install --dry-run --debug -f $(FILE) $(NAME) $(HELM_REPO)/$(HELM_CHART) --version $(HELM_VERSION)

delete:
	-helm delete $(NAME)

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHONY: port-forward port-forward-all
.port-forward:
	kubectl port-forward pods/$(chain)-genesis-0 $(localrpc):26657 &
	kubectl port-forward pods/$(chain)-genesis-0 $(localrest):1317 &

num_chains = $(shell yq -r ".chains | length" $(FILE))
port-forward-all:
	echo "Port forwarding all chains to localhost"
	for i in $(shell seq 0 $(num_chains)); do \
  		$(MAKE) .port-forward \
  			chain=$$(yq -r ".chains[$$i].name" $(FILE)) \
  			localrpc=$$(yq -r ".chains[$$i].ports.rpc" $(FILE)) \
  			localrest=$$(yq -r ".chains[$$i].ports.rest" $(FILE)); \
	done
	echo "Port forwarding explorer to localhost"
	kubectl port-forward service/explorer 8080:8080 &

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

.PHONY: check-forward
check-forward-all:
	while true ; do \
		for port in $$(yq -r "(.chains[] | .ports | .rpc, .rest), .explorer.localPorts.rest" $(FILE)); do \
			nc -vz 127.0.0.1 $$port; \
		done; \
		sleep 10; \
	done
