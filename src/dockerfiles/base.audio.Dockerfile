FROM base-maupod:latest
RUN DEBIAN_FRONTEND=noninteractive apt-get install -y ffmpeg mediainfo id3v2 flac
