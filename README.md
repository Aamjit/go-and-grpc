## Developement Roadmap

### Table of Contents

1. [Initialization](#initialization)
2. [Request/Response](#requestresponse-api)

#### Initialization

-   Initialize a Go project in root directory.

    ```
    go mod init github.com/your-username/your-project-name
    ```

-   Create three new directories in root.

    ```
    --> client
    --> proto
    --> server
    ```

-   Firstly, we will be using `GRPC` for this project to demonstrate and learn about the `GRPC` protocols and illustrate how we can use different functionalities of `GRPC` to establish communication between server and clients.

    -   Install the necessary packages for `GRPC` in the root directory or install it as an executable package in your system.

    we will use choco to fetch and install our `protobuf` package from its source.

    ```
    choco install protobuf
    ```

    After that we can check if it is properly installed in our system by running the following commands.

    ```
    protoc --version
    ```

-   Once installed, Create a `init.proto` file inside `/proto/` and initialize with the following code.

    ```
    syntax="proto3";
    option go_package = "./proto";
    package grpc_service;

    message NoParams{};
    message HelloRequest{
        string name = 1;
    };
    message HelloResponse{
        string message = 1;
    };
    message Lists {
        repeated string listItem = 1;
    }
    message MessageLists{
        repeated string message = 1;
    }

    service GrpcService{
        rpc GetHello(NoParams) returns (HelloResponse) {};
        rpc ServerStreaming(Lists) returns (stream HelloResponse);
        rpc ClientStreaming(stream HelloRequest) returns (MessageLists);
        rpc BidirectionalStreaming (stream HelloRequest) returns (stream HelloResponse);
    }
    ```

-   After this, we can generate our `.pb.go` file, which will be used as import package from our Go package and access the automatically generated functions by `protobuf`.

    ```
    protoc --go_out=. --go-grpc_out=. proto/init.proto
    ```

    If it is successful, it will execute without any error or message, and automatically generate 2 new files for us inside `proto` directory.

    ```
    -- init.pb.go
    -- init_grpc.pb.go
    ```

-   Create and Populate a new file: `/server/main.go`

    ```
    package main
    import (
        "log"
        "net"
    )

    type helloServer struct {
        pb.GrpcServiceServer
    }
    const (port = ":8080")

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

    ```

    Here, we initialize our `main.go` in our server component by creating a `TCP Listener` with port: `8080`.

    Create a new GRPC server and attached it to our proto service handler which we auto generated using protobuf using the register method for server.

-   Similarly, let us populate the `client/main.go`

    ```
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
    }
    ```

    Here, we initialize our `main.go` in our client component by creating a new GRPC
    client with the server address and port. We also attach the proto service handler
    to the client using the `NewGrpcServiceClient` method.

    Keep a note that we still haven't make use of the `client`, which is a GRPC client interface.

#### Request/Response API

-   Create a method to handle a simple Request/Response function between our server and client.

    -   Go to your main package in your server and add the function below

        ```
        func (s *helloServer) GetHello(ctx context.Context, req *pb.NoParams) (*pb.HelloResponse, error) {
            return &pb.HelloResponse{
                Message: "Hello",
            }, nil
        }
        ```

        Note, we will keep the function name same as we have maintained in our proto file.

    -   And in your main for client,

        ```
        func main() {
            conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
            if err != nil {
                log.Fatalf("Could not reach server %v", err)
            }

            defer conn.Close()

            client := pb.NewGrpcServiceClient(conn)

            callGetHello(client)
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
        ```

    -   Now, we can run these two programs separately in two terminals, and watch the results

        > server/main.go

        ![hello server](/assets/images/hello-server.png)

        > client/main.go

        ![hello client](/assets/images/hello-client.png)
