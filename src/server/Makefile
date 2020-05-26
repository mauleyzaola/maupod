build-dev:
	@echo "[INFO] building micro services in development mode"
	docker build -f ./artwork.Dockerfile -t maupod-artwork .

clean: down
	rm -f ./app

dev: build-dev up logs

down:
	docker-compose stop -f ./dev.docker-compose.yml
	docker-compose down -f ./dev.docker-compose.yml

logs:
	docker-compose logs -f

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
	docker-compose up -f ./dev.docker-compose.yml -d
