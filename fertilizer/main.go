package main

import (
	"github.com/Thomas-X/sandwhich-maker/fertilizer/fertilizerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func (s *server) GetFertilizer(ctx context.Context, in *fertilizerpb.GetFertilizerRequest) (*fertilizerpb.GetFertilizerResponse, error) {
	return &fertilizerpb.GetFertilizerResponse{
		Price: int32(float32(in.Amount) * float32(2.32)),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.FERTILIZER_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	fertilizerpb.RegisterFertilizerServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
