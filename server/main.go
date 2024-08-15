package main

import (
	"log"
	"net"

	pb "github.com/Aamjit/go-grpc/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GrpcServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGrpcServiceServer(grpcServer, &helloServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start GRPC server %v", err)
	}

	log.Printf("Server started at %v", lis.Addr())

}
