browser:
	$(MAKE) -C ../ui-player browser

build-dev: build-player
	@echo "[INFO] building micro services in development mode"

build-dockers: test
	@echo "[INFO] building micro services in production mode"
	bash -x ./scripts/build_docker.sh
	$(MAKE) -C . clean-binaries

build-player:
	@echo "[INFO] building player in host"
	go build -o maupod-player ./cmd/player

clean: down
	rm -f ./maupod-*

clean-binaries:
	rm -f maupod-*

clean-dockers: clean-binaries
	docker image ls mauleyzaola/maupod-* -q | xargs docker rmi

dev: build-dev up logs

dev-ui:
	$(MAKE) -C ../ui-player dev

down:
	docker-compose -f dev.docker-compose.yml stop
	docker-compose -f dev.docker-compose.yml down

logs:
	docker-compose -f dev.docker-compose.yml logs -f --tail 20

force:
prepare-proto: force
	rm -rf ./pkg/pb
	mkdir -p ./pkg/pb

prepare-proto-node: force
	rm -rf ./cmd/socket/nodepb
	mkdir -p ./cmd/socket/nodepb

proto: proto-node prepare-proto
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*

proto-node: prepare-proto-node
    # commonjs was the only way of make it work from nodejs with generated js files
	protoc --proto_path=proto --js_out=import_style=commonjs,binary:cmd/socket/nodepb ./proto/*

pg-backup:
	docker-compose -f dev.docker-compose.yml exec pg pg_dump maupod -U postgres > $(HOME)/Downloads/postgres.maupod.sql

psql:
	docker-compose -f dev.docker-compose.yml exec pg psql postgresql://postgres:nevermind@pg:5432/maupod

test:
	go test -cover ./...

up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose -f dev.docker-compose.yml up --build -d
