KEYS_FILE="../../charts/devnet/configs/keys.json"

PORT_PERSISTENCE=26657
PORT_GAIA=26653

CHAIN_ID_PERSISTENCE=core-1
CHAIN_ID_GAIA=gaia-4

setup-persistence-keys:
	CHAIN_BIN=persistenceCore CHAIN_ID=$(CHAIN_ID_PERSISTENCE) \
	DENOM=uxprt CHAIN_DATA_DIR=.persistenceCore NODE_PORT=$(PORT_PERSISTENCE) \
	VALIDATOR_CONFIG=$(KEYS_FILE) \
	bash -e setup_keys.sh

clean-persistence-keys:
	rm -rf ~/.persistenceCore/keyring-test/

setup-gaia-keys:
	CHAIN_BIN=gaiad CHAIN_ID=$(CHAIN_ID_GAIA) \
	DENOM=uatom CHAIN_DATA_DIR=.gaia NODE_PORT=$(PORT_GAIA) \
	VALIDATOR_CONFIG=$(KEYS_FILE) \
	bash -e setup_keys.sh

clean-gaia-keys:
	rm -rf ~/.gaia/keyring-test/

setup-keys: setup-persistence-keys setup-gaia-keys

clean-keys: clean-persistence-keys clean-gaia-keys

## Upgrade commands

pre-upgade:
	bash -e pre_upgrade.sh

.PHONY: upgrade
upgrade:
	CHAIN_ID=$(CHAIN_ID_PERSISTENCE) bash -e upgrade.sh

.PHONY: post-upgrade
post-upgrade:
	bash -e post_upgrade.sh

## Lscosmos testing
setup-lscosmos:
	bash -e setup_lscosmos.sh

test-lscosmos:
	bash -e test.sh
