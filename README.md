# Pluto

No code is needed to launch multiplayer game servers. Customize the config file and enjoy the server!

## SDK

* [Golang](https://github.com/0x9n0p/pluto/tree/dev/sdk)
* Unity (WIP, Help wanted)
* Unreal Engine (Contributor needed)

## Build & Run

Pluto gives you a single binary file, so you may clone the project and execute the command below.

```bash
go build -o pluto bin/main.go && ./pluto
```

The binary file does not take arguments, but you can pass several environment variables to it.

```bash
PLUTO_DEBUG=true PLUTO_HTTP_ADMIN=localhost:9630 ./pluto
```

## Configuration

## Architecture Overview
