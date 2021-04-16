protoc can do a bunch of languages, but not Go.
You need a plugin for that.
The latest version is v2, canonically: google.golang.org/protobuf (github: https://github.com/protocolbuffers/protobuf-go)

That github repo contains both things we need
- the protoc plugin which is the proto->Go compiler
- the protobuf runtime that we link against

## Installing
Install `protoc`

`protoc-gen-go`'s main package isn't in the root of that github repo, so if you `go get` it you're told there's no go files.
To get the compiler plugin binary: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

## Coding
Your code needs to link against the "SDK"
`go get google.golang.org/protobuf/proto`
