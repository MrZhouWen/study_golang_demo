package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"micro_koala/example/grpc_example/hello"
	"os"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()
	c := hello.NewHelloServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: name})
	if err != nil {
		log.Fatal("could not greet: %v", err)
		fmt.Println(err)
	}
	fmt.Println(r.Reply)
}
