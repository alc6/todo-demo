service=$1

# make sure the path exist
mkdir -p ${service}rpc

# generate go files from .proto files
protoc \
  --proto_path=${service} \
  --go_out=${service}rpc --go_opt=paths=source_relative \
  --go-grpc_out=${service}rpc --go-grpc_opt=paths=source_relative \
  ${service}/${service}.proto