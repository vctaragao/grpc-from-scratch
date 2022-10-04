package main

import (
	"context"
	"fmt"
	"log"

	"github.com/vctaragao/grpc-from-scratch/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewSayHelloClient(conn)

	req := &pb.Request{
		Name: "World",
	}

	res, err := client.HelloMessage(context.Background(), req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.GetStatus())
}
