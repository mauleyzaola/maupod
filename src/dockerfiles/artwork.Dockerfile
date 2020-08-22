FROM ubuntu:20.04
RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y locales ffmpeg id3v2 flac
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8
ENV LANG en_US.UTF-8 \
    LANGUAGE=en_US.UTF-8 \
    GOPATH="/go" \
    ZONEINFO="/go/src/github.com/mauleyzaola/maupod/src/zoneinfo.zip"
COPY . /go/src/github.com/mauleyzaola/maupod/src
COPY maupod-artwork /app
ENTRYPOINT ["/app/maupod-artwork"]