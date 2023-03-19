#!/bin/bash
. "$(dirname "$0")"/common.sh

# Script to copy pb.go files from bazel build folder to appropriate location.
# Bazel builds to bazel-bin/... folder, script copies them back to original folder where .proto is.

bazel build //proto/...

file_list=()
while IFS= read -d $'\0' -r file; do
    file_list=("${file_list[@]}" "$file")
done < <($findutil -L "$(bazel info bazel-bin)"/proto -type f -regextype sed -regex ".*pb\.\(gw\.\)\?go$" -print0)

arraylength=${#file_list[@]}
searchstring="Anmol1696/starship/"

# Copy pb.go files from bazel-bin to original folder where .proto is.
for ((i = 0; i < arraylength; i++)); do
    color "34" "$destination"
    destination=${file_list[i]#*$searchstring}
    echo "Source: ${file_list[i]} Destination: $destination"
    mkdir -p "$(dirname $destination)" && touch "$destination"
    cp -R -L "${file_list[i]}" "$destination"
    chmod 755 "$destination"
done
