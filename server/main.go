package main

import (
	"flag"
	pb "github.com/wdgk/grpc-sampler/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	addrFlag = flag.String("addr", ":5000", "Address host:post")
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("New Request: %v", in.String())
	return &pb.HelloReply{Message: "Hello, " + in.Name + "!"}, nil
}

func main() {
	lis, err := net.Listen("tcp", *addrFlag)

	if err != nil {
		log.Fatalf("boo")
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)

}
