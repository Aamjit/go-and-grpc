package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Aamjit/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not reach server %v", err)
	}

	defer conn.Close()

	client := pb.NewGrpcServiceClient(conn)

	lists := &pb.Lists{
		ListItem: []string{"Bucket", "Brush", "ToothBrush", "Comb"},
	}

	/*
		Simple GRPC function
	*/
	// callGetHello(client)

	/*
		Server streaming
	*/
	callServerStreaming(client, lists)

}

func callGetHello(client pb.GrpcServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	res, err := client.GetHello(ctx, &pb.NoParams{})

	if err != nil {
		log.Fatalf("Failed to called: %v", err)
	}

	log.Printf("%v", res)
}

func callServerStreaming(client pb.GrpcServiceClient, itemList *pb.Lists) {

	log.Printf("Stream started!")

	stream, err := client.ServerStreaming(context.Background(), itemList)

	if err != nil {
		log.Fatalf("Streaming failed: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while streaming: %v", err)
		}
		log.Println(message)
	}

	log.Println("Streaming finished!")

}
