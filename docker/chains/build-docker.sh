#!/bin/bash

DOCKER_REPO=${DOCKER_REPO:=anmol1696}
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

function image_tag_exists() {
  local image=$1
  local tag=$2

  return 1 # false

  # Check if tag is latest, return false if it is
  if [[ "$tag" == "latest" ]]; then
    return 1
  fi

  out=$(docker pull $image:$tag)
  if [[ $? -eq 0 ]]; then
    return 0
  else
    return 1
  fi
}

build_chain_tag() {
  local chain=$1
  local tag=$2
  local push_image=$3

  local base=$(yq -r ".chains[] | select(.name==\"$chain\") | .base" versions.yaml)

  # Push docker image, if feature flags set
  local buildx_args=""
  if [[ "$push_image" == "push" || "$push_image" == "push-only" ]]; then
    if image_tag_exists $DOCKER_REPO/$chain $tag; then
      color yellow "image $DOCKER_REPO/$chain:$tag already exists, skipping docker build"
      return 0
    fi
    color green "will pushing docker image $DOCKER_REPO/$chain:$tag"
    buildx_args="--push"
  fi

  color yellow "building docker image $DOCKER_REPO/$chain:$tag for chain $chain"
  for n in {1..3}; do
    docker buildx build \
      --platform linux/amd64 \
      -t "$DOCKER_REPO/$chain:$tag" \
      . -f Dockerfile \
      --build-arg BASE_IMAGE=$base \
      --build-arg VERSION=$tag \
      $buildx_args && break
    color red "failed to build docker image, retrying in 5 seconds, retry: $n"
    sleep 5
    if [[ "$n" == "3" ]]; then
      color red "failed to build docker image, exiting"
      exit 1
    fi
  done
}

build_all_tags() {
  local chain=$1
  local push_image=$2

  # Get all tags for the chain
  local tags=$(yq -r ".chains[] | select(.name==\"$chain\") | .tags[]" versions.yaml)

  # Build and push all tags
  for tag in $tags; do
    build_chain_tag $chain $tag $push_image
  done
}

build_all_chains() {
  local push_image=$1

  # Get all chains
  local chains=$(yq -r ".chains[].name" versions.yaml)

  # Build and push all chains
  for chain in $chains; do
    build_all_tags $chain $push_image
  done
}


while [ $# -gt 0 ]; do
  case "$1" in
    -c|--chain)
      CHAIN="$2"
      shift 2 # past argument=value
      ;;
    -t|--tag)
      TAG="$2"
      shift 2 # past argument
      ;;
    --push)
      PUSH="push"
      shift # past argument
      ;;
    --push-only)
      PUSH="push-only"
      shift # past argument
      ;;
    -*|--*)
      echo "Unknown option $1"
      exit 1
      ;;
    *)
      ;;
  esac
done


if [[ $CHAIN == "all" ]]; then
  build_all_chains $PUSH
  exit 0
fi

if [[ $TAG == "all" ]]; then
  build_all_tags $CHAIN $PUSH
  exit 0
fi

build_chain_tag $CHAIN $TAG $PUSH
