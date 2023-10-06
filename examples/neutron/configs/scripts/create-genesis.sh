#!/bin/bash

DENOM="${DENOM:=untrn}"
COINS="${COINS:=100000000000000000uosmo}"
CHAIN_ID="${CHAIN_ID:=osmosis}"
CHAIN_BIN="${CHAIN_BIN:=osmosisd}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"

set -eux

jq -r ".genesis[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN init $CHAIN_ID --chain-id $CHAIN_ID --recover

# Add genesis keys to the keyring and self delegate initial coins
echo "Adding key...." $(jq -r ".genesis[0].name" $KEYS_CONFIG)
jq -r ".genesis[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $(jq -r ".genesis[0].name" $KEYS_CONFIG) --recover --keyring-backend="test"
$CHAIN_BIN add-genesis-account $($CHAIN_BIN keys show -a $(jq -r .genesis[0].name $KEYS_CONFIG) --keyring-backend="test") $COINS --keyring-backend="test"

# Add relayer key to the keyring and self delegate initial coins
echo "Adding key...." $(jq -r ".relayers[0].name" $KEYS_CONFIG)
jq -r ".relayers[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $(jq -r ".relayers[0].name" $KEYS_CONFIG) --recover --keyring-backend="test"
$CHAIN_BIN add-genesis-account $($CHAIN_BIN keys show -a $(jq -r .relayers[0].name $KEYS_CONFIG) --keyring-backend="test") $COINS --keyring-backend="test"

ls $CHAIN_DIR/config


