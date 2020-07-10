package ports

//go:generate protoc --proto_path ./internal/proto/ --go_out=plugins=grpc,paths=source_relative:./internal/pb ./internal/proto/*
