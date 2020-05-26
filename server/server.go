package server

import (
	"fmt"
	pb "golang.org/x/bisq-grpc-gateway/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Simple Echo Server

type callService struct{}

func newCallServiceServer() pb.CallServiceServer {
	return new(callService)
}

func (m *callService) Call(ctx context.Context, cmd *pb.Params) (*pb.Response, error) {
	fmt.Printf("RPC call params: (%q)\n", cmd.Params)
	// header, ok := metadata.FromIncomingContext(ctx)
	// fmt.Printf("OK: %t\n HEADER: %s\n", ok, header)

	// Echo back the request fields in the response to show how it works.
	return &pb.Response{
		Result: cmd.Params,
	}, nil
}

// public functions are upper cased by convention
func RunServer(address string) error {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	server := grpc.NewServer()

	pb.RegisterCallServiceServer(server, newCallServiceServer())
	log.Printf("RPC server registered on %s\n", address)

	server.Serve(listen)
	return nil
}
