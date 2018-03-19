package main

import (
	prt "../proto"
	"google.golang.org/grpc"
	"fmt"
    "context"
    "net"
)

var port = ":5000"

type server struct{}

func main() {
	fmt.Println("Server is starting...")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	prt.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}

func (s *server) Reply(ctx context.Context, in *prt.ReqMessage) (*prt.RepMessage, error) {
	//fmt.Println("gRPC Reply")
	return &prt.RepMessage{Message: in.String()}, nil
}

