package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not reach server %v", err)
	}

	defer conn.Close()
	// client := pb.NewGrpcServiceClient(conn)

	// lists := &pb.Lists{
	// 	ListItem: []string{"Bucket", "Brush", "ToothBrush", "Comb"},
	// }

}
