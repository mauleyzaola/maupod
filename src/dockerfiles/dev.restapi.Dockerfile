FROM golang:1.14 as golang-dev

RUN ["go", "get", "-v", "github.com/githubnemo/CompileDaemon"]

FROM golang-dev
WORKDIR /go/src/github.com/mauleyzaola/maupod/src
ENTRYPOINT CompileDaemon -directory=. -log-prefix=false -build="go build -o maupod-restapi ./cmd/restapi" -command="./maupod-restapi" -exclude-dir=node_modules