package main

import (
	"flag"
	"golang.org/x/net/context"
	"log"
	"net"
	pb "proto/helloworld.proto"
)

type server struct{}

func (s *server) SayHello(contrxt context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

}
