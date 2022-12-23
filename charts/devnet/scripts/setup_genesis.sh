#!/bin/bash

DENOM="${DENOM:=uosmo}"
COINS="${COINS:=100000000000000000uosmo}"
CHAIN_ID="${CHAIN_ID:=osmosis}"
CHAIN_BIN="${CHAIN_BIN:=osmosisd}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"

set -eu

jq -r ".genesis[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN init $CHAIN_ID --chain-id $CHAIN_ID --recover

# Add keys to keyring and self deletegate inital coins
for type in $(jq -r ". | keys[]" $KEYS_CONFIG)
do
  for ((i=0; i<$(jq -r ".$type | length" $KEYS_CONFIG); i++))
  do
    echo "Adding key...." $(jq -r ".$type[$i].name" $KEYS_CONFIG)
    jq -r ".$type[$i].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $(jq -r ".$type[$i].name" $KEYS_CONFIG) --recover --keyring-backend="test"
    $CHAIN_BIN add-genesis-account $($CHAIN_BIN keys show -a $(jq -r .$type[$i].name $KEYS_CONFIG) --keyring-backend="test") $COINS --keyring-backend="test"
  done
done

NUM_KEYS=$($CHAIN_BIN keys list --keyring-backend test --output json | jq -r ". | length")
echo "Number of keys added to keyring: $NUM_KEYS"

echo "Creating gentx..."
$CHAIN_BIN gentx $(jq -r ".genesis[0].name" $KEYS_CONFIG) 5000000000$DENOM --keyring-backend="test" --chain-id $CHAIN_ID

echo "Output of gentx"
cat $CHAIN_DIR/config/gentx/*.json | jq

echo "Running collect-gentxs"
$CHAIN_BIN collect-gentxs

ls $CHAIN_DIR/config

echo "Update genesis.json file with updated local params"
sed -i -e "s/\"stake\"/\"$DENOM\"/g" $CHAIN_DIR/config/genesis.json
sed -i "s/\"time_iota_ms\": \".*\"/\"time_iota_ms\": \"$TIME_IOTA_MS\"/" $CHAIN_DIR/config/genesis.json

jq -r '.app_state.staking.params.unbonding_time |= "90s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.slashing.params.downtime_jail_duration |= "6s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.deposit_params.max_deposit_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.deposit_params.min_deposit[0].amount |= "10"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.voting_params.voting_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.tally_params.quorum |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.tally_params.threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.gov.tally_params.veto_threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json

# Set wasm as permissioned or permissionless based on environment variable
wasm_permission="Nobody"
if [ $WASM_PERMISSIONLESS == "true" ]
then
  wasm_permission="Everybody"
fi

jq -r ".app_state.wasm.params.code_upload_access.permission |= \"${wasm_permission}\"" $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r ".app_state.wasm.params.instantiate_default_permission |= \"${wasm_permission}\"" $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json

$CHAIN_BIN tendermint show-node-id
