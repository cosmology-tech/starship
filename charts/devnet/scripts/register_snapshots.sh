#!/bin/bash

CHAIN_ID="${CHAIN_ID:=osmosis-1}"
VAL_NAME="${VAL_NAME:=osmosis}"
CHAIN_DIR="${CHAIN_DIR:=$HOME/.osmosisd}"
COLLECTOR_HOST="${COLLECTOR_HOST}"

set -euxo pipefail

snapshot_name=data_${VAL_NAME}_$(date "+%F-%H-%M-%S")

# Create the snapshot that will be uploaded
function create_snapshot {
  tar -czvf /opt/${snapshot_name}.tar.gz $CHAIN_DIR/data
}

# Register the snapshot to the collector service
function register_snapshot {
  url=${COLLECTOR_HOST}/chains/${CHAIN_ID}/validators/${VAL_NAME}/snapshots/${snapshot_name}.tar.gz
  curl -v -i ${url} -H'Content-Encoding: gzip' -H'Content-TYPE: application/gzip' --data-binary @/opt/${snapshot_name}.tar.gz
}

create_snapshot
register_snapshot