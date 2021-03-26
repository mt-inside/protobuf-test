package main

import (
	"context"
	"net"

	pb "github.com/mt-inside/protobuf-test/pkg/model"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBlackBookServer
}

func (s *server) GetContacts(ctx context.Context, req *pb.ContactsRequest) (*pb.AddressBook, error) {
	addrs := &pb.AddressBook{}
	addrs.People = append(addrs.People, &pb.Person{Name: "dave", Id: 0})
	addrs.People = append(addrs.People, &pb.Person{Name: "fred", Id: 1})

	return addrs, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterBlackBookServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
