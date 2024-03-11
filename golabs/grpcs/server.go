package main

import (
	"context"
	"fmt"
	"net"

	"github.com/sifer169966/protocol-benchmarks/golabs/grpcs/apis/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	serverHost := fmt.Sprintf("%s:%s", "", "8080")
	lis, err := net.Listen("tcp4", serverHost)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	hdl := newHandler()
	defer srv.Stop()
	pb.RegisterTestServer(srv, hdl)
	fmt.Println("serving grpc server at :8080")
	err = srv.Serve(lis)
	if err != nil {
		panic(err)
	}
}

var _ pb.TestServer = (*grpcHandler)(nil)

type grpcHandler struct {
	pb.UnimplementedTestServer
}

func newHandler() *grpcHandler {
	return &grpcHandler{}
}

func (hdl *grpcHandler) Test(_ context.Context, in *pb.Request) (*emptypb.Empty, error) {
	// log.Println("hi ", in.Name)
	return &emptypb.Empty{}, nil
}
