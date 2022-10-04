package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/vctaragao/grpc-from-scratch/pb"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedSayHelloServer
}

func (service Server) HelloMessage(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println("Hello", req.GetName())

	response := &pb.Response{
		Status: 200,
	}

	return response, nil
}

func main() {
	grpcServer := grpc.NewServer()

	pb.RegisterSayHelloServer(grpcServer, &Server{})

	port := ":5000"

	listener, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Servidor esperando requisições...")

	grpcError := grpcServer.Serve(listener)

	if grpcError != nil {
		log.Fatal(grpcError)
	}

}
