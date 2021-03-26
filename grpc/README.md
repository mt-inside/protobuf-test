First: see ../protobuf

## Recall
That protobuf github repo contains:
- the protoc plugin which is the proto->Go compiler
- the protobuf runtime that we link against

## gRPC
Similar to protobuf, we need two things:
- the protoc plugin which generates gRPC stubs (client and server), based on `rpc` directives in the `.proto` files
- the gRPC runtime which we link against

## Installing
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

This installs the stub generator.
It also depends on `grpc`, which causes that to be downloaded into `$GOPATH/src` ready for building against.

## Coding
Your code needs to import `google.golang.org/grpc`

## Building
gRPC is almost completely separate from protobuf.
You still need to give the options to get protoc to build go proto files.
You also give go-grpc options, which causes that, separate plugin to make a separate file with gRPC stubs in it (client & server).
Because the stubs are generated with the same go package, it's usual to emit them into the same dir as the proto files, and then including that package will get you both.

