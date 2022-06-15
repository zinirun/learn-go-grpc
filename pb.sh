protoc -I=. \                                                                                                                                 ✔  21:35:03  
            --go_out . --go_opt paths=source_relative \
            --go-grpc_out . --go-grpc_opt paths=source_relative \
            protos/v1/user/user.proto