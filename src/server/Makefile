build-dev:
	@echo "[INFO] building micro services in development mode"
	docker build -t maupod .
	cp ./maupod/.maupod.yml .

clean: down
	rm -f ./app

dev: build-dev up logs

down:
	docker-compose stop
	docker-compose down

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
	docker-compose up -d
