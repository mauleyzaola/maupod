FROM base-maupod:latest
ENV LANG=en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8 \
    GOPATH="/go" \
    ZONEINFO="/go/src/github.com/mauleyzaola/maupod/src/zoneinfo.zip"
COPY maupod-restapi /
COPY . /go/src/github.com/mauleyzaola/maupod/src
ENTRYPOINT ["/maupod-restapi"]