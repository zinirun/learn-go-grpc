proto_file_path=$1
if [ -z "$proto_file_path"]
    then echo "No argument: pass the proto file's path"
else
    protoc -I=. \
    --go_out . --go_opt paths=source_relative \
    --go-grpc_out . --go-grpc_opt paths=source_relative \
    $proto_file_path
fi