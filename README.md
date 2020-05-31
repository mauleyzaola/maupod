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