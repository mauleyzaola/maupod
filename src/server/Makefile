force:
prepare-proto: force
	rm -rf ./pkg/pb
	mkdir -p ./pkg/pb

proto: prepare-proto
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*

test:
	go test -cover ./...
