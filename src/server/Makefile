build-dev: build-player
	@echo "[INFO] building micro services in development mode"

build-player:
	@echo "[INFO] building player in host"
	go build -o maupod-player ./cmd/player

clean: down
	rm -f ./maupod-*

dev: build-dev up logs

down:
	docker-compose -f dev.docker-compose.yml stop
	docker-compose -f dev.docker-compose.yml down

logs:
	docker-compose -f dev.docker-compose.yml logs -f

force:
prepare-proto: force
	rm -rf ./cmd/socket/pb
	mkdir -p ./cmd/socket/pb

prepare-proto-node: force
	rm -rf ./cmd/socket/nodepb
	mkdir -p ./cmd/socket//nodepb

proto: prepare-proto
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*

proto-node: prepare-proto-node
	# protoc --proto_path=proto --js_out=./cmd/socket/nodepb,binary:. ./proto/*
    # commonjs was the only way of make it work from nodejs with generated js files
	protoc --proto_path=proto --js_out=import_style=commonjs,binary:cmd/socket/nodepb ./proto/*

test:
	go test -cover -count 1 ./...

up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose -f dev.docker-compose.yml up --build -d
