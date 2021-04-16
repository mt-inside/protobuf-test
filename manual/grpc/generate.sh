# Because we specify a fully-qualified go_package in the protos, these are both a bit degenerate
protoc \
    -I=./api \
    --go_out=pkg/model \
    --go_opt="paths=source_relative" \
    --go-grpc_out=pkg/model \
    --go-grpc_opt="paths=source_relative" \
    ./api/addressbook.proto 
