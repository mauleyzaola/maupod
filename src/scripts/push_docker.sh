#!/usr/bin/env bash

set -e

echo "$DOCKER_PASS" | docker login --username "$DOCKER_USER" --password-stdin

docker push mauleyzaola/maupod-restapi
docker push mauleyzaola/maupod-mediainfo
docker push mauleyzaola/maupod-audioscan
docker push mauleyzaola/maupod-artwork

