protoc -I=. \
--go_out . --go_opt paths=source_relative \
--go-grpc_out . --go-grpc_opt paths=source_relative \
protos/v2/user/user.proto