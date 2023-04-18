#!/bin/bash

# todo: change repo to cosmology
DOCKER_REPO=anmol1696
# Set default values for boolean arguments
PUSH=0
PUSH_LATEST=0

set -euo pipefail

function color() {
  local color=$1
  shift
  local black=30 red=31 green=32 yellow=33 blue=34 magenta=35 cyan=36 white=37
  local color_code=${!color:-$green}
  printf "\033[%sm%s\033[0m\n" "$color_code" "$*"
}

function is_directory {
    if [ -d "$1" ]; then
        return 0 # true
    else
        return 1 # false
    fi
}

docker_process_build() {
  local type=$1
  local process=$2
  local push_image=$3
  local push_latest=$4

  local tag=latest

  if ! is_directory "$type/$process"; then
    color red "$type/$process is not a valid directory, please make sure inputs are correct"
    exit 1
  fi

  echo "here...."
  if [[ "$type" == "chains" ]]; then
    local tag=$(cat $type/$process/Dockerfile | grep -oP 'CHAIN_VERSION=\$\{CHAIN_VERSION:-"\K[0-9.]+(?="})' | cut -d '"' -f1 | head -1)
  fi

  color yellow "building docker image $DOCKER_REPO/$process:$tag from file $type/$process/Dockerfile"

  docker buildx build --platform linux/amd64 -t "$DOCKER_REPO/$process:$tag" . -f $type/$process/Dockerfile
  echo "$DOCKER_REPO/$process:$tag"

  if [[ "$push_image" == "push" ]]; then
    color green "pushing docker image $DOCKER_REPO/$process:$tag"
    docker push "$DOCKER_REPO/$process:$tag"
  fi

  if [[ "$push_latest" == "latest" && "$type" == "chains" ]]; then
    color green "pushing docker image $DOCKER_REPO/$process:$tag as latest"
    docker tag "$DOCKER_REPO/$process:$tag" "$DOCKER_REPO/$process:latest"
    docker push "$DOCKER_REPO/$process:latest"
  fi
}

build_all_process() {
  local type=$1
  find "$type" -type f -print0 | while IFS= read -r -d '' file; do
    echo "Building for $type/$file"
    docker_process_build $type $file ${@:3}
  done
}

build_all_types() {
  find "." -type f -print0 | while IFS= read -r -d '' type; do
    echo "Building for $type"
    build_all_process $type "all" ${@:3}
  done
}

while [ $# -gt 0 ]; do
  case "$1" in
    -t|--type)
      TYPE="$2"
      shift 2 # past argument=value
      ;;
    -p|--process)
      PROCESS="$2"
      shift 2 # past argument
      ;;
    --push)
      PUSH="push"
      shift # past argument
      ;;
    --push-latest)
      PUSH_LATEST="latest"
      shift # past argument
      ;;
    -*|--*)
      echo "Unknown option $i"
      exit 1
      ;;
    *)
      ;;
  esac
done

if [[ $TYPE == "all" ]]; then
  build_all_types "all" "all" $PUSH $PUSH_LATEST
  exit 0
fi

if [[ $PROCESS == "all" ]]; then
  build_all_process $TYPE "all" $PUSH $PUSH_LATEST
  exit 0
fi

docker_process_build $TYPE $PROCESS $PUSH $PUSH_LATEST
