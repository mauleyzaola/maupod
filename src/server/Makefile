# adjust to your environment as needed
export MEDIA_STORE="/media/mau/music-library/music"

build:
	@echo "[INFO] building micro services in development mode"
	docker build -f ./dockerfiles/artwork.Dockerfile -t maupod-artwork .
	docker build -f ./dockerfiles/audioscan.Dockerfile -t maupod-audioscan .
	docker build -f ./dockerfiles/mediainfo.Dockerfile -t maupod-mediainfo .
	docker build -f ./dockerfiles/restapi.Dockerfile -t maupod-restapi .

clean: down
	rm -f ./maupod-*

dev: build up logs

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
	go test -cover -count 1 ./...

up:
	@echo "[INFO] starting docker-compose micro services"
	docker-compose up -d
