package server

import (
	"context"

	pb "github.com/zscrub/rchat_server/protos"
)

func (s *RouteServer) SendMessage(ctx context.Context, message *pb.Message) (*pb.Response, error) {
	// store message

	return &pb.Response{Detail: "hey lol"}, nil
}
