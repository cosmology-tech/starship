#!/bin/bash

DENOM="${DENOM:=uosmo}"
COINS="${COINS:=100000000000000000uosmo}"
CHAIN_ID="${CHAIN_ID:=osmosis}"
CHAIN_BIN="${CHAIN_BIN:=osmosisd}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"
KEYS_CONFIG="${KEYS_CONFIG:=configs/keys.json}"
WASM_PERMISSIONLESS="${WASM_PERMISSIONLESS:=true}"

set -eu

# check if the binary has genesis subcommand or not, if not, set CHAIN_GENESIS_CMD to empty
CHAIN_GENESIS_CMD=$($CHAIN_BIN 2>&1 | grep -q "genesis-related subcommands" && echo "genesis" || echo "")

jq -r ".genesis[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN init $CHAIN_ID --chain-id $CHAIN_ID --recover

# Add genesis keys to the keyring and self delegate initial coins
echo "Adding key...." $(jq -r ".genesis[0].name" $KEYS_CONFIG)
jq -r ".genesis[0].mnemonic" $KEYS_CONFIG | $CHAIN_BIN keys add $(jq -r ".genesis[0].name" $KEYS_CONFIG) --recover --keyring-backend="test"
$CHAIN_BIN $CHAIN_GENESIS_CMD add-genesis-account $($CHAIN_BIN keys show -a $(jq -r .genesis[0].name $KEYS_CONFIG) --keyring-backend="test") $COINS --keyring-backend="test"

echo "Creating gentx..."
$CHAIN_BIN $CHAIN_GENESIS_CMD gentx $(jq -r ".genesis[0].name" $KEYS_CONFIG) 5000000000$DENOM --keyring-backend="test" --chain-id $CHAIN_ID

echo "Output of gentx"
cat $CHAIN_DIR/config/gentx/*.json | jq

echo "Running collect-gentxs"
$CHAIN_BIN $CHAIN_GENESIS_CMD collect-gentxs

ls $CHAIN_DIR/config

echo "Update genesis.json file with updated local params"
sed -i -e "s/\"stake\"/\"$DENOM\"/g" $CHAIN_DIR/config/genesis.json
sed -i "s/\"time_iota_ms\": \".*\"/\"time_iota_ms\": \"$TIME_IOTA_MS\"/" $CHAIN_DIR/config/genesis.json

jq -r '.app_state.staking.params.unbonding_time |= "200s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r '.app_state.slashing.params.downtime_jail_duration |= "6s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json

# overide for non 47 sdk
if [ "$(jq -r '.app_state.gov.params' $CHAIN_DIR/config/genesis.json)" == "null" ]; then
    jq -r '.app_state.gov.deposit_params.max_deposit_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.deposit_params.min_deposit[0].amount |= "10"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.voting_params.voting_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.tally_params.quorum |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.tally_params.threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.tally_params.veto_threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
else
    # overrides for 47 sdk
    jq -r '.app_state.gov.params.max_deposit_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.params.min_deposit[0].amount |= "10"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.params.voting_period |= "30s"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.params.quorum |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.params.threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
    jq -r '.app_state.gov.params.veto_threshold |= "0.000000000000000000"' $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
fi
# Set wasm as permissioned or permissionless based on environment variable
wasm_permission="Nobody"
if [ $WASM_PERMISSIONLESS == "true" ]
then
  wasm_permission="Everybody"
fi

jq -r ".app_state.wasm.params.code_upload_access.permission |= \"${wasm_permission}\"" $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
jq -r ".app_state.wasm.params.instantiate_default_permission |= \"${wasm_permission}\"" $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json

if [ "$CHAIN_BIN" == "meshd" ]
then
  jq -r ".app_state.meshsecurity.params.epoch_length |= \"10\"" $CHAIN_DIR/config/genesis.json > /tmp/genesis.json; mv /tmp/genesis.json $CHAIN_DIR/config/genesis.json
fi

$CHAIN_BIN tendermint show-node-id
