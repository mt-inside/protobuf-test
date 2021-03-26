package main

import (
	"context"
	"time"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/mt-inside/protobuf-test/pkg/model"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := pb.NewBlackBookClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetContacts(ctx, &pb.ContactsRequest{})
	if err != nil {
		panic(err)
	}

	spew.Dump(resp)
}
