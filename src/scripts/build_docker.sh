#!/usr/bin/env bash

set -e

prefix="./cmd/"

for dir in ./cmd/*; do
    if [[ -f $dir ]]; then
        echo "[INFO] $dir is not a directory, skipping..."
        continue
    fi

    name=${dir#"$prefix"}
    dockerfile="dockerfiles/$name.Dockerfile"

    # check the docker file exist for this directory
    if [[ ! -f  "$dockerfile" ]]; then
        echo "[WARNING] cannot stat $name.Dockerfile, skipping..."
        continue
    fi

    echo "[INFO] building $name"

#     buildLinuxApp "$name"
#     buildDocker "$name" "$dockerfile"
#     cleanUp "$name"
#     cd ../
done

# remove unused docker images
# docker image prune -f
