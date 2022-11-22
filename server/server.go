package server

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/zscrub/rchat_server/protos"
	"google.golang.org/grpc"
)

func InitServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	if err != nil {
		log.Fatal("Server failed to listen: ", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRChatServer(grpcServer, &ChatServer{})
	grpcServer.Serve(lis)
}
