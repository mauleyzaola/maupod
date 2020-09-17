os_name = $(shell uname -s | tr A-Z a-z)
pwd = $(shell pwd)
xopen =
ifeq ($(os_name),linux)
	xopen="xdg-open"
else ifeq ($(os_name),darwin)
	xopen="open"
endif

browser:
	$(MAKE) -C ../ui-player browser

clean: down

clean-binaries:
	rm -f maupod-*

clean-dockers: clean-binaries
	docker image prune -f
	docker image ls mauleyzaola/maupod-* -q | xargs docker rmi

dev: dev-build dev-up dev-logs

dev-build:
	go build -o maupod-player ./cmd/player

dev-clean: clean-binaries dev-down

dev-down:
	docker-compose -f ./dev.docker-compose.yml stop
	docker-compose -f ./dev.docker-compose.yml down

dev-logs:
	docker-compose -f ./dev.docker-compose.yml logs -f --tail 20

dev-up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose -f ./dev.docker-compose.yml up --build -d

dev-ui:
	$(MAKE) -C ../ui-player dev

dockers: test
	@echo "[INFO] building micro services in production mode"
	docker build -f ./dockerfiles/base.Dockerfile . --tag "base-maupod:latest"
	docker build -f ./dockerfiles/base.audio.Dockerfile . --tag "base-maupod-audio:latest"
	bash -x ./scripts/build_docker.sh
	$(MAKE) -C . clean-binaries

down:
	docker-compose stop
	docker-compose down --remove-orphans

force:

logs:
	docker-compose logs -f

prepare-proto: force
	sudo rm -rf ./pkg/pb
	mkdir -p ./pkg/pb

prepare-proto-node: force
	rm -rf ./cmd/socket/nodepb
	mkdir -p ./cmd/socket/nodepb

prod: dockers
	bash -x ./scripts/push_docker.sh

proto: proto-node prepare-proto
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*

proto-node: prepare-proto-node
    # commonjs was the only way of make it work from nodejs with generated js files
	protoc --proto_path=proto --js_out=import_style=commonjs,binary:cmd/socket/nodepb ./proto/*

pg-backup:
	docker-compose exec pg pg_dump maupod -U postgres > $(HOME)/Downloads/postgres.maupod.sql

pg-restore: clean
	bash -x ./scripts/pg_restore.sh

psql:
	docker-compose -f ./dev.docker-compose.yml exec pg psql postgresql://postgres:nevermind@pg:5432/maupod

server: dev-build up logs

swagger: force
	docker-compose -f swagger.docker-compose.yml up -d
	sleep 2
	$(xopen) "http://localhost:8888"

up:
	docker-compose up -d --build

update: clean
	git pull
	bash -x ./scripts/pull_docker.sh

test:
	go test -cover ./...
