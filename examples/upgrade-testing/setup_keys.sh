#!/bin/bash

CHAIN_BIN="${CHAIN_BIN:=persistenceCore}"
DENOM="${DENOM:=uxprt}"
CHAIN_DATA_DIR="${CHAIN_DATA_DIR:=.persistenceCore}"
CHAIN_ID="${CHAIN_ID:=persistencecore-1}"
NODE_HOST="${NODE_HOST:=localhost}"
NODE_PORT="${NODE_PORT:=26657}"

set -eu

jq -r ".genesis[0].mnemonic" $VALIDATOR_CONFIG | $CHAIN_BIN keys add $(jq -r ".genesis[0].name" $VALIDATOR_CONFIG) --recover --keyring-backend="test"

# Add keys to keyringg
for ((i=0; i<$(jq -r '.validators | length' $VALIDATOR_CONFIG); i++))
do
  jq -r ".validators[$i].mnemonic" $VALIDATOR_CONFIG | $CHAIN_BIN keys add $(jq -r ".validators[$i].name" $VALIDATOR_CONFIG) --recover --keyring-backend="test"
done

echo "Update client.toml file"
sed -i -e 's#keyring-backend = ".*"#keyring-backend = "test"#g' $HOME/$CHAIN_DATA_DIR/config/client.toml
sed -i -e 's#output = ".*"#output = "json"#g' $HOME/$CHAIN_DATA_DIR/config/client.toml
sed -i -e 's#broadcast-mode = ".*"#broadcast-mode = "block"#g' $HOME/$CHAIN_DATA_DIR/config/client.toml
sed -i -e "s#chain-id = \".*\"#chain-id = \"$CHAIN_ID\"#g" $HOME/$CHAIN_DATA_DIR/config/client.toml
sed -i -e "s#node = \".*\"#node = \"tcp://$NODE_HOST:$NODE_PORT\"#g" $HOME/$CHAIN_DATA_DIR/config/client.toml

$CHAIN_BIN status 2>&1
