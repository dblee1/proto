package main

import (
	"fmt"
	"flag"
	"context"
	"time"

	pb "github.com/dblee1/proto"
	
	"google.golang.org/grpc"
)

const (
	port = "localhost:50051"
	Msg = "Hello World!"
)

func main() {
	fmt.Println("gRPC client start!")

	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect gRPC server: ", err)
	}
	defer conn.Close()

	c := pb.NewHiClient(conn)

	name := Msg
	flag.Parse()
	if flag.NArg() > 0{
		name = flag.Arg(0)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	r, err := c.SayHi(ctx, &pb.HiRequest{Name: name})
	if err != nil {
		fmt.Println("could not greet: ", err)
	}
	fmt.Println("Greeting from the Server: ", r.Message)

	r, err = c.CountLength(ctx, &pb.HiRequest{Name: name})
        if err != nil {
		fmt.Println("could not count: ",err)
        }
        fmt.Println("Counting from the Server: ", r.Message)
}
