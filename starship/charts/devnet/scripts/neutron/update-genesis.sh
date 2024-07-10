#!/bin/bash

DENOM="${DENOM:=uosmo}"
CHAIN_BIN="${CHAIN_BIN:=osmosisd}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"

set -eux

ls $CHAIN_DIR/config

echo "Update genesis.json file with updated local params"
sed -i "s/\"time_iota_ms\": \".*\"/\"time_iota_ms\": \"$TIME_IOTA_MS\"/" $CHAIN_DIR/config/genesis.json

sed -i -e "s/\"denom\": \"stake\",/\"denom\": \"$DENOM\",/g" "$GENESIS_FILE"
sed -i -e "s/\"mint_denom\": \"stake\",/\"mint_denom\": \"$DENOM\",/g" "$GENESIS_FILE"
sed -i -e "s/\"bond_denom\": \"stake\"/\"bond_denom\": \"$DENOM\"/g" "$GENESIS_FILE"

$CHAIN_BIN tendermint show-node-id
