package main

import (
	"context"
	"log"
	"net"
	"time"

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

	log.Printf("Server started at %v", lis.Addr())

	grpcServer := grpc.NewServer()

	pb.RegisterGrpcServiceServer(grpcServer, &helloServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start GRPC server %v", err)
	}

}

func (s *helloServer) GetHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

func (apx *helloServer) ServerStreaming(req *pb.Lists, stream pb.GrpcService_ServerStreamingServer) error {

	log.Printf("Request: %v", req.ListItem)

	for _, listItem := range req.ListItem {
		log.Printf("Sending item: %v", listItem)

		res := &pb.HelloResponse{
			Message: "This is: " + listItem,
		}

		if err := stream.Send(res); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	return nil
}
