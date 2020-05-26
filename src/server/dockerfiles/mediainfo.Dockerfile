FROM golang:1 as golang-dev

RUN apt-get update && apt-get install -y mediainfo
RUN ["go", "get", "-v", "github.com/githubnemo/CompileDaemon"]

FROM golang-dev
WORKDIR /go/src/github.com/mauleyzaola/maupod/src/server
ENTRYPOINT CompileDaemon -directory=. -log-prefix=false -build="go build -o maupod-mediainfo ./maupod/cmd/mediainfo" -command="./maupod-mediainfo"