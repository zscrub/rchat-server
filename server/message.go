package server

import (
	"context"
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
		// key := serialize(in.Content)
		key := in.Content

		s.mu.Lock()
		s.message[key] = append(s.message[key], in)

		rn := make([]*pb.Message, len(s.message[key]))
		copy(rn, s.message[key])
		s.mu.Unlock()

		for _, note := range s.message[key] {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
