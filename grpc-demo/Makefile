PROTOC := protoc
PROTO_FILES := $(shell find models -type f -iname "*.proto")

protobuf:
	${PROTOC} --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	${PROTO_FILES}
