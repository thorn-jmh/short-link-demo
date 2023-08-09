# Short-Link-Demo

## How to run it

1. first create database
2. complete `comfig.yaml`. (see `config_example.yaml`)
3. run it using docker or just build it.

### Docker

```shell
docker build . -t short_link_demo
docker run -d --name short_link_demo short_link_demo
```

### Go build

make sure you have go>=1.20.5

```shell
go run .
```



## Test

run `go test ./internal/controller`



## CodeGen

run `go generate ./internal/dao/ent`
