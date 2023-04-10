generate-grpc:
	@protoc --go_out=api/data --go_opt=paths=source_relative --go-grpc_out=api/data --go-grpc_opt=paths=source_relative -I api/protos api/protos/data.proto