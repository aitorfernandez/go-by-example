package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/status"
)

var addrs = []string{":50051", ":50052", ":50053"}

type pbServer struct {
	addr string
}

func (s *pbServer) UnaryEcho(ctx context.Context, req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{Message: fmt.Sprintf("%s (from %s)", req.GetMessage(), s.addr)}, nil
}
func (s *pbServer) ServerStreamingEcho(*pb.EchoRequest, pb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *pbServer) ClientStreamingEcho(pb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}
func (s *pbServer) BidirectionalStreamingEcho(pb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "not implemented")
}

func startServer(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen %w", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &pbServer{addr: addr})

	log.Println("serving on", addr)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %w", err)
	}
}

func main() {
	var wg sync.WaitGroup
	for _, addr := range addrs {
		wg.Add(1)
		go func(addr string) {
			defer wg.Done()
			startServer(addr)
		}(addr)
	}
	wg.Wait()
}
