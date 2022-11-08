EXAMPLE = chain
FILE = custom-values.yaml

HELM_REPO = shuttle
HELM_CHART = devnet

setup:
	helm repo add $(HELM_REPO) https://anmol1696.github.io/shuttle/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART)

install:
	helm install -f $(EXAMPLE)/$(FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

debug:
	helm install --dry-run --debug -f $(EXAMPLE)/$(FILE) $(EXAMPLE) ../charts/$(HELM_CHART)

install-local:
	$(MAKE) install FILE=local-values.yaml

install-custom:
	$(MAKE) install FILE=custom-values.yaml

delete:
	-helm delete $(EXAMPLE)

###############################################################################
###                              Port forward                              ###
###############################################################################

.PHONY: port-forward port-forward-all
.port-forward:
	kubectl port-forward pods/$(chain)-genesis-0 $(localrpc):26657 &
	kubectl port-forward pods/$(chain)-genesis-0 $(localrest):1317 &

num_chains = $(shell yq -r ".chains | length" $(EXAMPLE)/$(FILE))
port-forward-all:
	echo "Port forwarding all chains to localhost"
	for i in $(shell seq 0 $(num_chains)); do \
  		$(MAKE) .port-forward \
  			chain=$$(yq -r ".chains[$$i].name" $(EXAMPLE)/$(FILE)) \
  			localrpc=$$(yq -r ".chains[$$i].ports.rpc" $(EXAMPLE)/$(FILE)) \
  			localrest=$$(yq -r ".chains[$$i].ports.rest" $(EXAMPLE)/$(FILE)); \
	done
	echo "Port forwarding explorer to localhost"
	kubectl port-forward service/explorer 8080:8080 &

.PHONY: stop-forward
stop-forward:
	-pkill -f "port-forward"

.PHONY: check-forward
check-forward-all:
	while true ; do \
		for port in $$(yq -r "(.chains[] | .ports | .rpc, .rest), .explorer.localPorts.rest" $(EXAMPLE)/$(FILE)); do \
			nc -vz 127.0.0.1 $$port; \
		done; \
		sleep 10; \
	done
