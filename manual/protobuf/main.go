package main

import (
	"os"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/davecgh/go-spew/spew"
	pb "github.com/mt-inside/protobuf-test/pkg/model"
)

func main() {
	addrs := &pb.AddressBook{}
	addrs.People = append(addrs.People, &pb.Person{Name: "dave", Id: 0})
	addrs.People = append(addrs.People, &pb.Person{Name: "fred", Id: 1})

	out, err := proto.Marshal(addrs)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile("/tmp/proto", out, 0644)
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	in, err := os.ReadFile("/tmp/proto")
	if err != nil {
		panic(err)
	}

	addrsRead := &pb.AddressBook{}
	err = proto.Unmarshal(in, addrsRead)
	if err != nil {
		panic(err)
	}

	spew.Dump(addrsRead)
}
