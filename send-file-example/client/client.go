package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	pb "github.com/dblee1/proto/practice/proto"

	"google.golang.org/grpc"
)

const (
	port       = "localhost:50051"
	maxMsgSize = 107374182400
)

func main() {
	fmt.Println("gRPC client start!")
/*
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMsgSize)))
*/

	conn, err := grpc.Dial(port, /*grpc.WithDefaultCallOptions(grpc.MaxCallSendMsgSize(maxMsgSize))*/ grpc.WithInsecure())
	if err != nil {
		fmt.Println("failed to connect gRPC server: ", err)
	}
	defer conn.Close()

	c := pb.NewFileClient(conn)

	// setting timeout
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// get file list
	path := "/home/dblee/source"
	files, err := GetFileList(path)
	if err != nil {
		fmt.Println("get File list error")
		panic(err)
	}

	if len(files) == 0 {
		panic("no files in directory")
	}

	// get file data
	for _, file := range files {
		dat, err := ReadFile(file.Name())
		if err != nil {
			fmt.Printf("cannot read file, file name: %v, %v\n", file.Name(), err)
			continue
		}

		// send file data to server
		r, err := c.SendFile(ctx, &pb.FileRequest{Filename: file.Name(), Data: dat})
		if err != nil {
			fmt.Printf("failed to send file, file name: %v, %v\n", file.Name(), r.Response)
		}
		fmt.Printf("Send file data to Server, filename: %v, %v\n", file.Name(), r.Response)
	}

}

func ReadFile(filename string) ([]byte, error) {
	path := "/home/dblee1/source/" + filename

	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("failed to reaf file: ", filename)
		return nil, err
	}

	return dat, nil
}

func GetFileList(dirname string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir("../../../../source")
	if err != nil {
		fmt.Println("failed to read directory: ", err)
		return nil, err
	}
	return files, nil
}
