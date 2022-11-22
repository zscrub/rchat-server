package server

import (
	"context"
	"io"

	pb "github.com/zscrub/rchat_server/protos"
)

func (s *RouteServer) SendMessage(ctx context.Context, message *pb.Message) (*pb.Response, error) {
	// store message
	return &pb.Response{Detail: "Connected"}, nil
}

func (s *RouteServer) ChatSession(stream pb.RChat_ChatSessionServer) error {
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
		s.routeNotes[key] = append(s.routeNotes[key], in)

		rn := make([]*pb.RouteNote, len(s.routeNotes[key]))
		copy(rn, s.routeNotes[key])
		s.mu.Unlock()

		for _, note := range s.routeNotes[key] {
			if err := stream.Send(note); err != nil {
				return err
			}
		}
	}
}
