#!/bin/bash

set -eux

DENOM="${DENOM:=uatom}"
COINS="${COINS:=100000000000000000uatom}"
CHAIN_ID="${CHAIN_ID:=cosmoshub-1}"
CHAIN_BIN="${CHAIN_BIN:=gaiad}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.gaia}"
NODE_URL="${NODE_URL:=http://0.0.0.0:26657}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"

PROPOSAL_FILE="${PROPOSAL_FILE}"
KEY_NAME="ics-setup"

add_key() {
  # Add test keys to the keyring and self delegate initial coins
  echo "Adding key...." $(jq -r ".test[0].name" $KEYS_CONFIG)
  jq -r ".test[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $KEY_NAME --recover --keyring-backend="test"
  echo $($CHAIN_BIN keys show -a $KEY_NAME --keyring-backend="test")
}

submit_proposal() {
  echo "Get all porposals"
  $CHAIN_BIN query gov proposals --output json --node $NODE_URL
  echo "Submit gov proposal on chain"
  $CHAIN_BIN tx gov submit-proposal consumer-addition $PROPOSAL_FILE \
    --from $KEY_NAME \
    --chain-id $CHAIN_ID \
    --node $NODE_URL \
    --keyring-backend="test" \
    --output json \
    --yes
}
