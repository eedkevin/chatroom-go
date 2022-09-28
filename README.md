# chatroom-demo

## Features 
- [x] multiple chat rooms
- [x] multiple participants per chat room
- [x] chat room history
- [x] participants should be able to see the 10 latest messages before joining a chat room
- [x] new chat room messages must be transmitted to other participants in real-time

## Prerequisite
- Docker Engine 20.10.8 or above
- Docker Compose v2.0.0-rc.1 or above

## Quick Start
1. Lift up the application using `make`
```
$ make compose-up
```

or `docker compose`

```
$ docker compose up --build
```

2. Trial from web browser
```
$ open http://localhost:8080
```

check `Makefile` for more shortcuts to ease local development