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

Set the following environment variables in your `~/.bashrc` or `~/.bash_profile`

```
export MAUPOD_BASE_IP_ADDRESS=192.168.0.135
export MAUPOD_SOCKET_PORT=8181
export MAUPOD_MEDIA_STORE=/mnt/music-library
export MAUPOD_ARTWORK="$MAUPOD_MEDIA_STORE/artwork"
export REACT_APP_MAUPOD_API="http://$MAUPOD_BASE_IP_ADDRESS:7400"
export REACT_APP_MAUPOD_ARTWORK="http://$MAUPOD_BASE_IP_ADDRESS:7401"
export REACT_APP_MAUPOD_SOCKET="ws://$MAUPOD_BASE_IP_ADDRESS:$MAUPOD_SOCKET_PORT"
export HOST="$MAUPOD_BASE_IP_ADDRESS"
```

In this example my ip address is `192.168.0.135` and my music library is located at this path: `/mnt/music-library`. Change these variables to match your environment

It is important you keep the other values unchanged as above, otherwise maupod won't work


In one terminal go to `src/` directory in the repo

```
make dev
```

Once that is done, run in another terminal
```
make dev-ui
```

And finally, on a third terminal, run this

```
./maupod-player
```

Browser should automatically start at http://192.168.0.135:3000 (whatever your ip address was defined above)


### Development

Ensure to meet the following requirements

[Tagger](src/pkg/taggers/README.md)

[Protoc](src/docs/protocol-buffers.md)

[Mediainfo](src/docs/mediainfo.md)

### Related Software

These packages will make your life easier, although not mandatory

[Flacon](src/docs/flacon.md)

[Kid3](src/docs/kid3.md)

#### Running Postgres Isolated

```
docker run --rm --name postgres-maupod -v $HOME/data/maupod/pg/data:/var/lib/postgresql/data -p 5432:5432 -e POSTGRES_PASSWORD=nevermind -d postgres:9.5
```