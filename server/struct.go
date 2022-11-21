package server

import (
	pb "github.com/zscrub/rchat_server/protos"
)

type RouteServer struct {
	pb.UnimplementedRChatServer
}
