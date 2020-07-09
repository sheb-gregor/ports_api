get_proto:
	protoc --proto_path ./proto/ --go_out=paths=source_relative:./pkg/prototypes ./proto/*

