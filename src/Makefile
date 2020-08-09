build-dev: build-player
	@echo "[INFO] building micro services in development mode"

build-player:
	@echo "[INFO] building player in host"
	go build -o maupod-player ./cmd/player

clean: down
	rm -f ./maupod-*

dev: build-dev up logs

dev-ui:
	# this value should be in the environment instead
	$(MAKE) -C ../ui-player dev

down:
	docker-compose stop
	docker-compose down

logs:
	docker-compose logs -f

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
	docker-compose exec pg pg_dump maupod -U postgres > $(HOME)/Downloads/postgres.maupod.sql

psql:
	docker-compose exec pg psql postgresql://postgres:nevermind@pg:5432/maupod

test:
	go test -cover -count 1 ./...

up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose up --build -d
