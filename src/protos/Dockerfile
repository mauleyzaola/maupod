ARG GO_VERSION=1.14.9
ARG PROTOC_VERSION=3.13.0
ARG PROTOC_GEN_GO_VERSION=1.25.0

FROM golang:${GO_VERSION}

# Install necessary dependencies
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool

# Install protoc
ARG PROTOC_VERSION
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
RUN unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local bin/protoc
RUN unzip -o protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local 'include/*'
RUN rm -f protoc-${PROTOC_VERSION}-linux-x86_64.zip

# Install protoc-gen-grpc-gateway and protoc-go-inject-tag
RUN go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
RUN go get -u github.com/favadi/protoc-go-inject-tag

# Install protoc-gen-go
ARG PROTOC_GEN_GO_VERSION
RUN go get -v google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go get -v google.golang.org/protobuf/cmd/... && \
    cd $GOPATH/src/google.golang.org/protobuf/cmd/protoc-gen-go/ && \
    git checkout v${PROTOC_GEN_GO_VERSION} && \
    go install && \
    cd $GOPATH/src/google.golang.org/grpc/cmd/protoc-gen-go-grpc && \
    go install

# Create work folder
WORKDIR /go/src/cloudproto