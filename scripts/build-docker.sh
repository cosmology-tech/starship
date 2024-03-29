#!/bin/bash

# todo: change repo to cosmology
DOCKER_REPO=${DOCKER_REPO:=anmol1696}
# Set default values for boolean arguments
PUSH=0
PUSH_LATEST=0
FORCE=0

DOCKER_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )/../docker"

set -euo pipefail

function color() {
  local color=$1
  shift
  local black=30 red=31 green=32 yellow=33 blue=34 magenta=35 cyan=36 white=37
  local color_code=${!color:-$green}
  printf "\033[%sm%s\033[0m\n" "$color_code" "$*"
}

function set_docker_buildx() {
  set +e
  out=$(docker buildx create --name chain-builder --use > /dev/null 2>&1)
  if [[ $? -ne 0 ]]; then
    docker buildx use chain-builder
  fi
  set -e
}

function is_directory {
    if [ -d "$1" ]; then
        return 0 # true
    else
        return 1 # false
    fi
}

function image_tag_exists() {
  local image=$1
  local tag=$2

  # Check if tag is latest, return false if it is
  if [[ "$tag" == "latest" ]]; then
    return 1 # false
  fi

  out=$(docker pull $image:$tag)
  if [[ $? -eq 0 ]]; then
    return 0 # true
  else
    return 1 # false
  fi
}

docker_process_build() {
  local type=$1
  local process=$2
  local version=$3
  local push_image=$4
  local push_latest=$5

  if [[ "$type" == "chains" ]]; then
    color red "we dont build docker image anymore for chains from this script"
    exit 1
  fi

  local base=""
  if [ -f "$DOCKER_DIR/$type/$process/versions.yaml" ]; then
    base=$(yq -r ".base" $DOCKER_DIR/$type/$process/versions.yaml)
  fi
  local tag=${version##*/}

  if ! is_directory "$DOCKER_DIR/$type/$process"; then
    color red "$DOCKER_DIR/$type/$process is not a valid directory, please make sure inputs are correct"
    exit 1
  fi

  if [[ "$FORCE" -ne 1 ]]; then
    if image_tag_exists $DOCKER_REPO/$process $tag; then
      color yellow "image $DOCKER_REPO/$process:$tag already exists, skipping docker build"
      return 0
    fi
    color green "image not found remote, will build docker image $DOCKER_REPO/$process:$tag"
  fi

  # Push docker image, if feature flags set
  local buildx_args=""
  if [[ "$push_image" == "push" || "$push_image" == "push-only" ]]; then
    color green "will pushing docker image $DOCKER_REPO/$process:$tag"
    buildx_args="--push"
  fi

  # Build docker image if push-only is not set
  if [[ "$push_image" != "push-only" ]]; then
    color yellow "building docker image $DOCKER_REPO/$process:$tag from file $DOCKER_DIR/$type/$process/Dockerfile"
    for n in {1..3}; do
      docker buildx build \
        --platform linux/amd64,linux/arm64 \
        -t "$DOCKER_REPO/$process:$tag" . \
        --build-arg BASE_IMAGE=$base \
        --build-arg VERSION=$version \
        -f $DOCKER_DIR/$type/$process/Dockerfile \
        $buildx_args && break
      color red "failed to build docker image, retrying in 5 seconds, retry: $n"
      sleep 5
      if [[ "$n" == "3" ]]; then
        color red "failed to build docker image, exiting"
        exit 1
      fi
    done
    echo "$DOCKER_REPO/$process:$tag"
  fi

  # Push the docker image with tag as latest
  if [[ "$push_latest" == "latest" ]]; then
    color green "pushing docker image $DOCKER_REPO/$process:$tag as latest"
    docker tag "$DOCKER_REPO/$process:$tag" "$DOCKER_REPO/$process:latest"
    docker push "$DOCKER_REPO/$process:latest"
  fi
}

build_all_versions() {
  local type=$1
  local process=$2
  local versions=(latest)
  if [ -f "$DOCKER_DIR/$type/$process/versions.yaml" ]; then
    versions=$(yq -r ".versions[]" $DOCKER_DIR/$type/$process/versions.yaml)
  fi
  for version in $versions; do
    echo "Building for $type/$process:$version"
    docker_process_build $type $process $version ${@:4}
  done
}

build_all_process() {
  local type=$1
  for process in $DOCKER_DIR/$type/*/; do
    process="${process%*/}"
    process="${process##*/}"
    echo "Building for $type/$process"
    build_all_versions $type $process "all" ${@:4}
  done
}

build_all_types() {
  for type in $DOCKER_DIR/*/; do
    type="${type%*/}"
    type="${type##*/}"
    if [[ "$type" == "chains" ]]; then
      color red "we dont build docker image anymore for chains from this script"
      continue
    fi
    echo "Building for all $type"
    build_all_process $type "all" "all" ${@:4}
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
    -v|--version)
      VERSION="$2"
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
    --push-latest)
      PUSH_LATEST="latest"
      shift # past argument
      ;;
    --force)
      FORCE=1
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

set_docker_buildx

if [[ $TYPE == "all" ]]; then
  build_all_types "all" "all" "all" $PUSH $PUSH_LATEST
  exit 0
fi

if [[ $PROCESS == "all" ]]; then
  build_all_process $TYPE "all" "all" $PUSH $PUSH_LATEST
  exit 0
fi

if [[ $VERSION == "all" ]]; then
  build_all_versions $TYPE $PROCESS "all" $PUSH $PUSH_LATEST
  exit 0
fi

docker_process_build $TYPE $PROCESS $VERSION $PUSH $PUSH_LATEST
