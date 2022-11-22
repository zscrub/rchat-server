package server

import (
	"sync"

	pb "github.com/zscrub/rchat_server/protos"
)

type RouteServer struct {
	pb.UnimplementedRChatServer

	mu         sync.Mutex // protects routeNotes
	routeNotes map[string][]*pb.RouteNote
}
