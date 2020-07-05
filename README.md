# maupod

These are a set of applications that eventually should accomplish:

- [x] Automatic Media Management
- [ ] Web UI for media streaming
- [ ] Remote control for media player

## Requirements

* Docker
* NodeJS

[mpv](src/docs/mpv.md)

## Running

For the time being, only in development mode. There are no releases yet, 
so the procedure is kind of tricky

This environment variable needs to point to the directory where your media files live

```
export MEDIA_STORE="/media/mau/music-library/music"
```

And this other needs to point to the ip of the backend

```
export REACT_APP_API_URL="http://localhost:8000"
```

In one terminal 

```
make dev
```

Once that is done, run in another terminal
```
make dev-ui
```

Browser will automatically start at http://localhost:3000

#### Artwork

If you want to enable this feature, you'll need to configure a valid directory in the `src/server/docker-compose.yml` file in the `maupod-restapi` service

Default is `~/Downloads/artwork`

```
- $HOME/Downloads/artwork:/artwork
```

### Development

Ensure to meet the following requirements

[Tagger](src/pkg/taggers/README.md)

[Protoc](src/docs/protocol-buffers.md)

[Mediainfo](src/docs/mediainfo.md)

### Related Software

These packages will make your life easier, although not mandatory

[Flacon](src/docs/flacon.md)

[Kid3](src/docs/kid3.md)