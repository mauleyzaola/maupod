FROM base-maupod:latest
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ffmpeg mediainfo id3v2 flac
ENV LANG en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8 \
    GOPATH="/go" \
    ZONEINFO="/go/src/github.com/mauleyzaola/maupod/src/zoneinfo.zip"
