EXAMPLE = mesh-security
FILE = local-values.yaml

HELM_REPO = shuttle
HELM_CHART = devnet

setup:
	helm repo add $(HELM_REPO) https://anmol1696.github.io/shuttle/
	helm repo update
	helm search repo $(HELM_REPO)/$(HELM_CHART)

install:
	helm install -f $(EXAMPLE)/$(FILE) $(EXAMPLE) $(HELM_REPO)/$(HELM_CHART)

install-local:
	$(MAKE) install FILE=local-values.yaml

install-custom:
	$(MAKE) install FILE=custom-values.yaml

delete:
	-helm delete $(EXAMPLE)
