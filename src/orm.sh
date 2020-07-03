#!/usr/bin/env bash


echo "[INFO] building database model"
echo "[INFO] building database model"
if [[ ! $(which sqlboiler) ]]; then
    echo "[INFO] sqlboiler not found, go to github.com/mauleyzaola/sqlboiler and execute ./install.sh"
  exit 1
fi

if [[ ! $(which sql-migrate) ]]; then
    echo "[INFO] sql-migrate not found, run the command below outside of GOPATH"
    echo "go get -u github.com/rubenv/sql-migrate/..."
    exit 1
fi

echo "[INFO] re-generating orm files"
sqlboiler --wipe --no-tests --output ./pkg/dbdata/orm --pkgname orm --no-auto-timestamps --no-hooks --struct-tag-casing snake psql
