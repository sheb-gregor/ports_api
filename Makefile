get_proto:
	protoc --proto_path ./internal/proto/ --go_out=plugins=grpc,paths=source_relative:./internal/pb ./internal/proto/*

build_docker_client_api:
	docker build -t client_api --build-arg SERVICE=client_api --build-arg CONFIG=tmpl .

build_docker_port_domain_service:
	docker build -t port_domain_service --build-arg SERVICE=port_domain_service --build-arg CONFIG=tmpl .


build_docker: build_docker_client_api build_docker_port_domain_service

test:
	go test ./...

lint:
	golangci-lint run --out-format tab
