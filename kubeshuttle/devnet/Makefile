# Set helm args based on only file to show
HELM_NAME = kubeshuttle
ifneq ($(FILE),)
HELM_ARGS += --show-only $(FILE)
endif

KEYS_CONFIG = configs/keys.json

all: delete install

debug:
	helm template --dry-run --debug --generate-name . $(HELM_ARGS)

install:
	helm install --replace --debug $(HELM_NAME) . $(HELM_ARGS)

upgrade:
	helm upgrade --debug $(HELM_NAME) . $(HELM_ARGS)

delete:
	helm delete --debug $(HELM_NAME)

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
