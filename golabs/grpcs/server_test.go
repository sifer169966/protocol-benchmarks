package main

import (
	"context"
	"testing"

	"github.com/sifer169966/protocol-benchmarks/golabs/grpcs/apis/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func BenchmarkTest100(b *testing.B) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("127.0.0.1:8080", opts...)
	if err != nil {
		b.Fatal(err)
	}
	client := pb.NewTestClient(conn)
	defer conn.Close()

	for i := 0; i < b.N; i++ {
		requestgRPC100(client, b)
	}
}

func requestgRPC100(c pb.TestClient, b *testing.B) {
	req := &pb.Request{Name: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"}
	_, err := c.Test(context.Background(), req)
	if err != nil {
		b.Fatal(err)
	}
}
