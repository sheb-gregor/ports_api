# Ports API

### Requirements 

- go 1.13+ 
- mockgen — `go get github.com/golang/mock/mockgen`
- forge — `go get -u github.com/lancer-kit/forge`
- protoc - https://developers.google.com/protocol-buffers
- protoc-gen-go - `go get github.com/golang/protobuf/protoc-gen-go`
- golangci-lint — https://github.com/golangci/golangci-lint
- easyjson - `go get -u github.com/mailru/easyjson/...`


### Build and Running

Project consist from two applications(services):

1. `client_api` — provides REST API and data parser
1. `port_domain_service` — provides gRPC and data storage

- Build apps with Docker

```shell script
# client_api:
docker build -t client_api --build-arg SERVICE=client_api --build-arg CONFIG=tmpl .

# port_domain_service:
docker build -t port_domain_service --build-arg SERVICE=port_domain_service --build-arg CONFIG=tmpl .

# OR

make build_docker
```

- Run 

```shell script
docker-compose up -d
```

### Configuration

Service configuration placed at [./env](/env) directory. 
Configuration consist from `.yaml` file for each app and `.env` file with secrets.

Naming convention for configuration files:

- `<env>.<app>_cfg.yaml`
- `<env>.env`

Where `<env>` is identifier of the deployment group/env, `<app>` is name of service.

### Compile app binary

```shell script
# client_api:
./build.sh client_api 

# port_domain_service:
./build.sh port_domain_service 
```


### Usage

Usage of client_api 

```text
Usage of ./client_api:
  -config string
        path to config file (default "config.yaml")
```

Usage of port_domain_service 

```text
Usage of ./port_domain_service:
  -config string
        path to config file (default "config.yaml")
```


### Testing

```shell script
go test ./...
```


### client_api

Client API exposes 3 HTTP routes:

- `GET /info` - returns build & version of application;

Response:

```json
{
  "name": "ports_client_api",
  "version": "1.1.0",
  "build": "37dce1e-dirty.",
  "tag": "master"
}
```

- `GET /ports` - returns list of `Ports` with pagination;

Request params:

| name     | in    | type    | values                                                                       |
|----------|-------|---------|------------------------------------------------------------------------------|
| order    | query | string  | asc, desc                                                                    |
| page     | query | integer | 0..total                                                                     |
| pageSize | query | integer | 1..500                                                                       |
| orderBy  | query | string  | unlocode, name, city, country, timezone, code, extra, created_at, updated_at |


Response:

> `GET http://127.0.0.1:8080/port?order=asc&page=1&pageSize=2&orderBy=name`

```json
{
  "page": 1,
  "pageSize": 2,
  "order": "asc",
  "order_by": "name",
  "total": 816,
  "records": [
    {
      "name": "Aalborg",
      "city": "Aalborg",
      "country": "Denmark",
      "coordinates": [
        9.921747,
        57.04882
      ],
      "province": "North Denmark Region",
      "timezone": "Europe/Copenhagen",
      "unlocs": [
        "DKAAL"
      ],
      "code": "40903"
    },
    {
      "name": "Aarhus",
      "city": "Aarhus",
      "country": "Denmark",
      "coordinates": [
        10.22,
        56.15
      ],
      "province": "Midtjylland",
      "timezone": "Europe/Copenhagen",
      "unlocs": [
        "DKAAR"
      ],
      "code": "40906"
    }
  ]
}
```

- `GET /ports/{unlocode}` - returns singe `Port` record;

Response:

```json
{
  "name": "Abu Dhabi",
  "city": "Abu Dhabi",
  "country": "United Arab Emirates",
  "coordinates": [
    54.37,
    24.47
  ],
  "province": "Abu Z¸aby [Abu Dhabi]",
  "timezone": "Asia/Dubai",
  "unlocs": [
    "AEAUH"
  ],
  "code": "52001"
}
```

