package main

import (
	"github.com/aitorfernandez/go-by-example/grpc/builder"
	"google.golang.org/grpc/resolver"
)

func main() {

}

func init() {
	var addrs = []string{"localhost:50051", "localhost:50052", "localhost:50053"}
	resolver.Register(builder.New("test", addrs))
}
