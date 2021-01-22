package main

import (
	"fmt"
	"context"
	"net"
	"io/ioutil"

	pb "github.com/dblee1/proto/practice/proto"

	"google.golang.org/grpc"
)

const port = ":50051"

type grpcServer struct{}

func (s *grpcServer) SendFile(ctx context.Context, in *pb.FileRequest) (*pb.FileResponse, error) {
	// if server receive data
	fmt.Printf("Received: %v, File Size: %v\n", in.Filename, len(string(in.Data)))

	// store file to directory
	path := "/home/dblee1/destination/"
	filename := in.Filename
	err := ioutil.WriteFile(path+filename, in.Data, 0644)
	if err != nil {
		fmt.Printf("failed to write file, file name: %v, %v\n", filename, err)
		response := "NOK"
		return &pb.FileResponse{Response: response}, err
	}

	// send response to client
	response := "OK"
	fmt.Printf("Send Back to client: %s\n", response)
	return &pb.FileResponse{Response: response}, nil
}

func main (){
	fmt.Printf("gRPC Server Start at port %v\n", port)

	ret, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	s := grpc.NewServer(grpc.MaxRecvMsgSize(107374182400), grpc.MaxSendMsgSize(107374182400))
	pb.RegisterFileServer(s, &grpcServer{})
	if err := s.Serve(ret); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
