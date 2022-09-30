# chatroom-demo

## Folder structure
```sh
.
├── cmd
│   └── app                     # cmd to start the service
├── internal                    # internal packages
│   ├── app                     # app
│   │   ├── adapter             # adapter layer
│   │   │   ├── http            # http adapter
│   │   │   │   ├── index       # index http adapter impl
│   │   │   │   ├── room        # room http adapter impl
│   │   │   │   └── user        # user http adapter impl
│   │   │   ├── repository      # repository impl
│   │   │   ├── service         # service impl
│   │   │   └── ws              # websocket adapter
│   │   ├── application         # application layer
│   │   │   ├── service         # service interface
│   │   │   └── usecase         # core business logic
│   │   ├── domain              # domain layer
│   │   │   ├── repository      # repository interface
│   │   │   └── vo              # value object - domain entity segments
│   │   └── infrastructure      # infrastructure layer
│   │       ├── inmemory        # inmemory
│   │       └── redis           # redis
│   └── cfg                     # config
├── pkg                         # public packages
│   └── client                  # websocket CLI client
├── public                      # frontend pages
├── testdata                    # test setup and mockup
│   └── fixture                 # test fixture
└── main.go                     # main entry
```

## Features 
- [x] multiple chat rooms
- [x] multiple participants per chat room
- [x] chat room history
  - when a user joined a chat room, historical messages will be deliveried
- [x] participants should be able to see the 10 latest messages before joining a chat room
  - a participant can view the 10 latest message via `GET /api/rooms/:id/thumbnail` endpoint before joining the chat room
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