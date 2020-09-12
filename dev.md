# Development Guide

This is a detailed guide to allow developing on Maupod

## Requirements

This is the list of software you will need to install at minimum, in order to run the Backend, UI and Player components

### Docker

#### Mac OS

https://hub.docker.com/editions/community/docker-ce-desktop-mac/

#### Ubuntu >= 18.04

```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu bionic stable"
sudo apt-get update && apt-cache policy docker-ce && sudo apt install docker-ce -y
sudo curl -L https://github.com/docker/compose/releases/download/1.22.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose
sudo usermod -aG docker ${USER}
su - ${USER}
```

### mpv

MPV is the program that plays the audio files and is required to be installed in your machine

Mac OS

```
brew update && brew install mpv
```

Ubuntu

```
sudo apt-get update && sudo apt-get install -y mpv
```

### Node / Yarn

Mac OS

```
brew update && brew install node yarn
```

Ubuntu

```
curl -sL https://deb.nodesource.com/setup_10.x | sudo bash -
sudo apt-get install -y nodejs
curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
sudo apt update && sudo apt install -y yarn
```

## UI

Covered above

## Backend

### Mac OS

```
brew update && brew install go postgresql
```

### Ubuntu

You can follow the instructions in the official site: https://golang.org/
or use this script at https://github.com/mauleyzaola/scripts/tree/master/go, which will
automatically download and configure golang for you

```
wget -O - https://raw.githubusercontent.com/mauleyzaola/scripts/master/go/go.install.sh | sh
```

For the make and postgres client, install these packages

```
sudo apt-get update && sudo apt-get install -y postgresql-client-common software-properties-common python g++ make cmake
```

###  Protobuf

Please stick to these versions so we don't check in a bunch of stuff in the repo

```
brew update && brew install protobuf
cd
GO111MODULE=off go get -v google.golang.org/grpc/cmd/protoc-gen-go-grpc
GO111MODULE=off go get -v google.golang.org/protobuf/cmd/...
cd $GOPATH/src/google.golang.org/protobuf/cmd/protoc-gen-go/
git checkout v1.25.0
go install
cd $GOPATH/src/google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install
```

You can check the right version using these commands
```
$ protoc-gen-go --version
protoc-gen-go v1.25.0

$ protoc --version
libprotoc 3.13.0
```

## Environment Variables

Maupod applications rely on environment in order to determine the ip address of the UI and Backend. Also to configure the root directory
for audio files and artwork

Make sure you have them configured as needed *before* starting the applications

```
export MAUPOD_BASE_IP_ADDRESS=192.168.0.135
export MAUPOD_SOCKET_PORT=8181
export MAUPOD_MEDIA_STORE=/mnt/music-library
export MAUPOD_ARTWORK="$MAUPOD_MEDIA_STORE/artwork"
export REACT_APP_MAUPOD_API="http://$MAUPOD_BASE_IP_ADDRESS:7400"
export REACT_APP_MAUPOD_ARTWORK="http://$MAUPOD_BASE_IP_ADDRESS:7401"
export REACT_APP_MAUPOD_SOCKET="ws://$MAUPOD_BASE_IP_ADDRESS:$MAUPOD_SOCKET_PORT"
export HOST="$MAUPOD_BASE_IP_ADDRESS"
```

Basically you would need to override these variables:
```
MAUPOD_BASE_IP_ADDRESS
MAUPOD_MEDIA_STORE
```

The rest are inferred automatically

Also, make sure you don't change the port numbers, otherwise docker won't be able to map them correctly

## Starting in Development Mode

Browse to `src/` direectory and run `make dev`. This command will build the docker images for developing the backend. Consider hot loading is
enabled, so each time you save files in golang, all the docker related containers will rebuild their binaries. This can take some time dependending on your machine

Browse to `src/` directory and run `make browser` if you're using a Desktop OS or run `make dev-ui` if you're using a Server OS (`dev-ui` won't start a browser, only the UI server)

UI is coded in React, and hot reloading feature works like backend does. This allows you to see the changes as you save automatically in the browser

If the browser didn't start automatically, you can go to `http://localhost:9990` or `http://your-ip-address:9990` to see the UI