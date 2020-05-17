#!/usr/bin/env bash


echo "[INFO] building database model"
if [[ ! $(which sqlboiler) ]]; then
    echo "[INFO] sqlboiler not found, trying to install now..."
    # need to build manually because of original sqlboiler package dependencies
    cd "$GOPATH/src/github.com/mauleyzaola"
    if ! git clone https://github.com/mauleyzaola/sqlboiler.git;then
      exit 1
    fi
    cd sqlboiler
    if ! go mod download;then
      exit 1
    fi
    cd drivers/sqlboiler-psql
    if ! go install;then
      exit 1
    fi
    cd ../../
    if ! go install;then
      exit 1
    fi
fi

if [[ ! $(which sql-migrate) ]]; then
    echo "[INFO] sql-migrate not found, run the command below outside of GOPATH"
    echo "go get -u github.com/rubenv/sql-migrate/..."
    exit 1
fi

echo "[INFO] re-generating orm files"
sqlboiler --wipe --no-tests --output ./pkg/data/orm --pkgname orm --no-auto-timestamps --no-hooks --struct-tag-casing snake psql
