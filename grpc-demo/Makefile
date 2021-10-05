PROTOC := protoc
PROTO_FILES := $(shell find models -type f -iname "*.proto")

protobuf:
	${PROTOC} -I ./models/vendor -I . \
	--go_out=./models/vendor --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out . \
	--grpc-gateway_opt logtostderr=true \
	--grpc-gateway_opt paths=source_relative \
	${PROTO_FILES}
