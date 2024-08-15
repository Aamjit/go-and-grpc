## Table of Contents

1. [Development Roadmap](#development-roadmap)

#### Development Roadmap

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
        rpc GetHello(HelloRequest) returns (HelloResponse) {};

        # GRPC method for server streaming
        rpc ServerStreaming(Lists) returns (stream HelloResponse);

        # GRPC method for client streaming
        rpc ClientStreaming(stream HelloRequest) returns (MessageLists);

        # GRPC method for to-from streaming
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

    const (
        port = ":8080"
    )

    func main() {
        lis, err := net.Listen("tcp", port)

        if err != nil {
            log.Fatalf("Failed to start server %v", err)
        }

    }

    ```

    Here, we initialize our `main.go` in our server component by creating a `TCP Listener` with port: `8080`.
