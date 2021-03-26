# source_relative stops the whole the go_package being appended to go_out
protoc \
    -I=./api \
    --go_out=pkg/model \
    --go_opt="paths=source_relative" \
    ./api/addressbook.proto 
