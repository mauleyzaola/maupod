FROM base-maupod:latest
ENV LANG en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8 \
    GOPATH="/go" \
    ZONEINFO="/go/src/github.com/mauleyzaola/maupod/src/zoneinfo.zip"
COPY . /go/src/github.com/mauleyzaola/maupod/src
COPY maupod-restapi /
COPY .maupod.yml /
ENTRYPOINT ["/maupod-restapi"]