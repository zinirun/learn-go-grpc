# .PHONY: generate-user-v2-gateway-proto
# generate-user-v2-gateway-proto:
    protoc -I . \
        --grpc-gateway_out . \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        protos/v2/user/user.proto