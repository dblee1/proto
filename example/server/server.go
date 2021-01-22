package main

import (
	"fmt"
	"context"
	"net"
	"strconv"

	pb "github.com/dblee1/proto"

	"google.golang.org/grpc"
)

const port = ":50051"

type server struct{}

func (s *server) SayHi(ctx context.Context, in *pb.HiRequest) (*pb.HiResponse, error) {
	fmt.Printf("Received: %v\n", in.Name)

	response := "Hi " + in.Name

	fmt.Printf("Send Back to client: %s\n", response)
	return &pb.HiResponse{Message: response}, nil
}

func (s *server) CountLength(ctx context.Context, in *pb.HiRequest) (*pb.HiResponse, error){
	fmt.Printf("Received: %v, Length: %d\n", in.Name, len(in.Name))

	msg := fmt.Sprintf("Count %v length: %v", in.Name, strconv.Itoa(len(in.Name)))
	fmt.Println("Send back to client: ", msg)

	return &pb.HiResponse{Message: msg}, nil
}

func main (){
	fmt.Printf("gRPC Server Start at port %v\n", port)

	ret, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterHiServer(s, &server{})
	if err := s.Serve(ret); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
