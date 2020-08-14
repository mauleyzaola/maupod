FROM golang:1.14 as golang-dev


RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y locales ffmpeg id3v2 flac libmagickwand-dev
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

ENV LANG en_US.UTF-8

RUN ["go", "get", "-v", "github.com/githubnemo/CompileDaemon"]

FROM golang-dev
WORKDIR /go/src/github.com/mauleyzaola/maupod/src
ENTRYPOINT CompileDaemon -directory=. -log-prefix=false -build="go build -o maupod-artwork ./cmd/artwork" -command="./maupod-artwork" -exclude-dir=node_modules