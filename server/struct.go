package server

import (
	"sync"

	pb "github.com/zscrub/rchat_server/protos"
)

type ChatServer struct {
	pb.UnimplementedRChatServer

	mu         sync.Mutex // protects routeNotes
	message map[string][]*pb.Message
}
