package server

import (
	"context"
	"fmt"
	"io"

	pb "github.com/zscrub/rchat_server/protos"
)

func (s *ChatServer) SendMessage(ctx context.Context, message *pb.Message) (*pb.Response, error) {
	// store message
	return &pb.Response{Detail: "Connected"}, nil
}

func (s *ChatServer) ChatSession(stream pb.RChat_ChatSessionServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		s.mu.Lock()
		s.message = append(s.message, in)
		fmt.Println(s.message)
		s.mu.Unlock()

		for _, note := range s.message {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
