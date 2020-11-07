SHELL := /bin/bash
WORKING_DIR := $(shell pwd)

docker-build:
	# Build docker image with all the dependencies to generate proto files
	docker build -t maupod-protogen .
	# Mount the local dir and run "make clean build inject-tag" inside container
	docker run --rm --name running-protogen -v ${WORKING_DIR}:/go/src/cloudproto maupod-protogen sh -c "make clean build inject-tag"
	
# Commands below are only used inside Docker container
build:
	$(foreach proto_file, $(shell find . -type f -name "*.proto"), \
		protoc -I ./ \
			-I $(shell dirname `which protoc`)/$(shell dirname $(shell readlink $(shell which protoc)))/../include \
			--go_out=paths=source_relative:. \
			--go-grpc_out=paths=source_relative:. \
			--grpc-gateway_out=paths=source_relative:. \
			$(proto_file); \
	)

inject-tag:
	$(foreach go_file, $(shell find . -type f -name "*.go"), \
		protoc-go-inject-tag -input=$(go_file); \
	)

clean:
	-rm -f $(shell find ./ -type f -name "*.go")
