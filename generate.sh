# Because we specify a fully-qualified go_package in the protos, these are both a bit degenerate
protoc \
    -I=./api \
    --go_out=. \
    --go_opt="module=github.com/mt-inside/protobuf-test" \
    ./api/addressbook.proto 
