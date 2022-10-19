#!/bin/bash

DENOM="${DENOM:=uosmo}"
CHAIN_BIN="${CHAIN_BIN:=osmosisd}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"
VAL_NAME="${VAL_NAME:=osmosis}"

set -eu

# Wait for the node to be synced
max_tries=10
while [[ $($CHAIN_BIN status 2>&1 | jq ".SyncInfo.catching_up") == true ]]
do
  if [[ max_tries -lt 0 ]]; then echo "Not able to sync with genesis node"; exit 1; fi
  echo "Still syncing... Sleeping for 15 secs. Tries left $max_tries"
  ((max_tries--))
  sleep 30
done

# Run create validator tx command
echo "Running txn for create-validator"
$CHAIN_BIN tx staking create-validator \
  --pubkey=$($CHAIN_BIN tendermint show-validator) \
  --moniker $VAL_NAME \
  --amount 5000000000$DENOM \
  --keyring-backend="test" \
  --chain-id $CHAIN_ID \
  --from $VAL_NAME \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000" \
  --fees 100000$DENOM \
  --gas="auto" \
  --gas-adjustment 1.5 --yes > /validator.log

cat /validator.log | jq
