# maupod

These are a set of applications that eventually should accomplish:

- [x] Automatic Media Management
- [ ] Web UI for media streaming
- [ ] Remote control for media player

## Requirements

* Docker
* NodeJS

## Running

This environment variable needs to point to the directory where your media files live

```
export MEDIA_STORE="/media/mau/music-library/music"
```

### Frontend

```
cd src/ui/ui-player
make dev
```

`ctrl+c` to stop

Start the browser at http://localhost:3000

### Backend

```
cd src/server
make dev
```

`ctrl+c` to stop then `make clean` to clean up docker stuff

#### Artwork

If you want to enable this feature, you'll need to configure a valid directory in the `src/server/docker-compose.yml` file in the `maupod-restapi` service

Default is `~/Downloads/artwork`

```
- $HOME/Downloads/artwork:/artwork
```