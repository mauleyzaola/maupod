#!/usr/bin/env bash

set -e

SOURCE="$HOME/Downloads/postgres.maupod.sql"
stat "$SOURCE"
docker run --name restore -d --rm -v $HOME/data/maupod/pg/data:/var/lib/postgresql/data postgres:9.5
sleep 5

docker cp $SOURCE restore:/postgres.maupod.sql
docker exec restore psql postgresql://postgres:nevermind@127.0.0.1:5432/postgres -c "drop database if exists maupod"
docker exec restore psql postgresql://postgres:nevermind@127.0.0.1:5432/postgres -c "create database maupod"
cat $SOURCE | docker exec -i restore psql postgresql://postgres:nevermind@127.0.0.1:5432/maupod
docker stop restore