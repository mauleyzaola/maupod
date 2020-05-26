build-dev:
	@echo "[INFO] building micro services in development mode"
	docker build -f ./artwork.Dockerfile -t maupod-artwork .
	docker build -f ./audioscan.Dockerfile -t maupod-audioscan .
	docker build -f ./restapi.Dockerfile -t maupod-restapi .

clean: down
	rm -f ./app

dev: build-dev up logs

down:
	docker-compose -f ./dev.docker-compose.yml stop
	docker-compose -f ./dev.docker-compose.yml down

logs:
	docker-compose -f ./dev.docker-compose.yml logs -f

force:
prepare-proto: force
	rm -rf ./pkg/pb
	mkdir -p ./pkg/pb

proto: prepare-proto
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*

test:
	go test -cover ./...

up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose -f ./dev.docker-compose.yml up -d
