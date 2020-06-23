package main

import (
	"context"
	"fmt"
	"log"
	"time"

	b "github.com/aitorfernandez/go-by-example/grpc/balancer"
	"github.com/aitorfernandez/go-by-example/grpc/builder"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	pb "google.golang.org/grpc/examples/features/proto/echo"
	"google.golang.org/grpc/resolver"
)

const (
	scheme      = "test"
	serviceName = "serviceName"
)

func callUnaryEcho(c pb.EchoClient, message string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.Message)
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	hwc := pb.NewEchoClient(cc)
	for i := 0; i < n; i++ {
		callUnaryEcho(hwc, "load_balancing")
	}
}

func main() {
	awesome, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", scheme, serviceName),
		grpc.WithBalancerName("my_awesome_balancer"),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer awesome.Close()

	makeRPCs(awesome, 3)
}

func init() {
	var addrs = []string{"localhost:50051", "localhost:50052", "localhost:50053"}
	resolver.Register(builder.New(scheme, addrs))

	balancer.Register(b.NewBuilder())
}
