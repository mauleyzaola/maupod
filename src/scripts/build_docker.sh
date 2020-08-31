#!/usr/bin/env bash

set -e

buildLinuxApp() {
    dir=$1
    appname=$2

    cd "$dir"
    echo "[INFO] building executable file: $appname"
    GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o "../../maupod-$appname" -ldflags "-X main.Version=$VERSION -X main.Hash=$HASH -X main.HashDate=$HASH_DATE" ./
    echo "[INFO] application:$appname has been built successfully"
    cd -
}

buildDocker(){
   echo "[INFO] building docker image: $1:$HASH using Dockerfile: $1.Dockerfile"
   docker build --force-rm -f "$2" . --tag "mauleyzaola/maupod-$1:latest"
}

VERSION=$(git tag | tail -n 1)
HASH=$(git log -1 --format="%H")/
HASH_DATE=$(git show --no-patch --pretty='%cd' --date=format:'%Y-%m-%dT%H:%M:%S')
VERSION="$VERSION"

prefix="./cmd/"

for dir in ./cmd/*; do
    if [[ -f $dir ]]; then
        continue
    fi

    name=${dir#"$prefix"}
    dockerfile="dockerfiles/$name.Dockerfile"

    # check the docker file exist for this directory
    if [[ ! -f  "$dockerfile" ]]; then
        continue
    fi

    echo "[INFO] building $name"

    buildLinuxApp "$dir" "$name"
    buildDocker "$name" "$dockerfile"
#     cd ../
done

# remove unused docker images

docker image prune -f