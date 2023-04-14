#!/bin/bash

function color() {
  local color=$1
  shift
  local black=30 red=31 green=32 yellow=33 blue=34 magenta=35 cyan=36 white=37
  local color_code=${!color:-$green}
  printf "\033[%sm%s\033[0m\n" "$color_code" "$*"
}

system=""
case "$OSTYPE" in
darwin*) system="darwin" ;;
linux*) system="linux" ;;
msys*) system="windows" ;;
cygwin*) system="windows" ;;
*) exit 1 ;;
esac
readonly system

# Get locations of pb.go and pb.gw.go files.
findutil="find"
# On OSX `find` is not GNU find compatible, so require "findutils" package.
if [ "$system" == "darwin" ]; then
    if [[ ! -x "/usr/local/bin/gfind" && ! -x "/opt/homebrew/bin/gfind" ]]; then
        color red "Make sure that GNU 'findutils' package is installed: brew install findutils"
        exit 1
    else
        export findutil="gfind"  # skipcq: SH-2034
    fi
fi

# Script to copy pb.go files from bazel build folder to appropriate location.
# Bazel builds to bazel-bin/... folder, script copies them back to original folder where .proto is.

bazel build //proto/...

file_list=()
while IFS= read -d $'\0' -r file; do
    file_list=("${file_list[@]}" "$file")
done < <($findutil -L "$(bazel info bazel-bin)"/proto -type f -regextype sed -regex ".*pb\.\(gw\.\)\?go$" -print0)

arraylength=${#file_list[@]}
searchstring="cosmology-tech/starship/"

# Copy pb.go files from bazel-bin to original folder where .proto is.
for ((i = 0; i < arraylength; i++)); do
    color blue "$destination"
    destination=${file_list[i]#*$searchstring}
    echo "Source: ${file_list[i]} Destination: $destination"
    mkdir -p "$(dirname $destination)" && touch "$destination"
    cp -R -L "${file_list[i]}" "$destination"
    chmod 755 "$destination"
done
