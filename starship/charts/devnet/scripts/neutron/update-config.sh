#!/bin/bash

STAKEDENOM=${DENOM:-untrn}
CHAIN_ID="${CHAIN_ID:=osmosis}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"

set -eux

ls $CHAIN_DIR

echo "Already done"
