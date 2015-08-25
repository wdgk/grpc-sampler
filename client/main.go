package main

import (
	"flag"
	pb "github.com/wdgk/grpc-sampler/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

var (
	addrFlag = flag.String("addr", "localhost:5000", "server address host:post")
)

func main() {
	conn, err := grpc.Dial(*addrFlag)

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "wdgk"})
	if err != nil {
		log.Fatalf("RPC error: %v", err)
	}
	log.Printf("Greeting: %s", resp.Message)
}
