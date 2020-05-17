proto:
	rm -rf ./pkg/pb
	mkdir -p ./pkg/pb
	protoc --proto_path=proto --go_out=pkg/pb --go_opt=paths=source_relative ./proto/*