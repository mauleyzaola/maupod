# Protocol Buffers

Maupod applications use `protobuf` for serializing/deserializing 
data across services

## Building

```
make proto
```

Files should be generated in `./pkg/pb` directory

## Installation

First, we need the `proto` program installed, next the go plugin

### Ubuntu

```
mkdir -p ~/bin && cd $_
sudo apt-get install autoconf automake libtool curl make g++ unzip -y
git clone https://github.com/protocolbuffers/protobuf.git
cd protobuf
git submodule update --init --recursive
./autogen.sh
./configure
make
make check
sudo make install
sudo ldconfig
```

### Mac OS

```
brew update && brew install protobuf
```

### Go Plugin

As for the go plugin we can install it directly in the `$GOBIN` path

```
cd
go get -v google.golang.org/protobuf/cmd/...
cd $GOPATH/src/google.golang.org/protobuf/cmd/protoc-gen-go/
go install
```