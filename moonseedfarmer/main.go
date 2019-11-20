package main

import (
	"github.com/Thomas-X/sandwhich-maker/moonseedfarmer/moonseedfarmerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type server struct {
}

func (s *server) GetMoonSeed(ctx context.Context, in *moonseedfarmerpb.GetMoonSeedRequest) (*moonseedfarmerpb.GetMoonSeedResponse, error) {
	price := 0
	switch in.BreadName {
	case "Soleil":
		price = 1
	}

	return &moonseedfarmerpb.GetMoonSeedResponse{
		Price: int32(price),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.MOONSEED_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	moonseedfarmerpb.RegisterMoonSeedFarmerServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
