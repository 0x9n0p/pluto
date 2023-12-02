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

The binary file does not take arguments, but you can pass environment variables to it.

```bash
PLUTO_HOST=localhost,127.0.0.1 PLUTO_HTTP_SERVER=0.0.0.0:80 PLUTO_PANEL_STORAGE=./tmp/ PLUTO_DEBUG=true ./pluto
```

## Panel

### APIs Examples

#### Create/Save pipelines

This API Creates/Saves the pipeline and returns the saved pipeline in response.

```bash
curl -X POST http://panel.localhost/api/v1/pipelines \                                                                                     
  -H 'Content-Type: application/json' \
  -d '{"name":"PIPELINE_NAME","processors":[{"name":"Write to input/output","arguments":[{"name":"io_interface","type":"InternalInterface","value":"ARGUMENT_VALUE"}]}]}'
```

```json
{
  "name": "PIPELINE_NAME",
  "processors": [
    {
      "name": "Write to input/output",
      "arguments": [
        {
          "name": "io_interface",
          "type": "InternalInterface",
          "value": "ARGUMENT_VALUE"
        }
      ]
    }
  ]
}
```

## Configuration

## Architecture Overview
