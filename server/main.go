package main

import (
	pb "github.com/alisoltanics/gRPC-go/add"
	"context"
	"net"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Add(ctx context.Context, request *pb.Request) (*pb.Response, error) {
	a, b := request.GetA(), request.GetB()
	result := a + b
	return &pb.Response{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterAddServiceServer(s, &server{})

	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
