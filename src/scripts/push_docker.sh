#!/usr/bin/env bash

set -e

docker login --username "mauleyzaola" --password "$DOCKER_PASS"

docker push mauleyzaola/maupod-restapi
docker push mauleyzaola/maupod-mediainfo
docker push mauleyzaola/maupod-audioscan
docker push mauleyzaola/maupod-artwork

