package main

import (
	"context"
	"github.com/Thomas-X/sandwhich-maker/fertilizer/fertilizerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"github.com/Thomas-X/sandwhich-maker/tomato/tomatopb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func setupFertilizer() (*fertilizerpb.FertilizerServiceClient, **grpc.ClientConn) {
	fertilizerConnPort := strconv.Itoa(shared.FERTILIZER_PORT)
	conn, err := grpc.Dial("localhost:"+fertilizerConnPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to fertilizer server %v", err)
	}
	c := fertilizerpb.NewFertilizerServiceClient(conn)
	log.Printf("Created fertilizer client %v", c)
	return &c, &conn
}

func GetFertilizerForTomato(amount int32, name string, c *fertilizerpb.FertilizerServiceClient) *tomatopb.Tomato {
	lettuce, err := (*c).GetFertilizer(context.Background(), &fertilizerpb.GetFertilizerRequest{
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("Couldn't get fertilizer for tomato: %v", err)
	}

	return &tomatopb.Tomato{
		Name:  name,
		Price: lettuce.Price,
	}
}


func (s *server) GetTomato(ctx context.Context, in *tomatopb.GetTomatoRequest) (*tomatopb.GetTomatoResponse, error) {
	var tomatoes []*tomatopb.Tomato
	c, conn := setupFertilizer()
	defer (*conn).Close()

	for i := 0; i < len(in.Tomatoes); i++ {
		switch in.Tomatoes[i] {
		case "dutch":
			tomatoes = append(tomatoes, GetFertilizerForTomato(1, "dutch", c))
		case "kumato":
			tomatoes = append(tomatoes, GetFertilizerForTomato(3, "kumato", c))
		case "roma":
			tomatoes = append(tomatoes, GetFertilizerForTomato(5, "roma", c))
		}
	}
	return &tomatopb.GetTomatoResponse{Tomatoes: tomatoes}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.TOMATO_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	tomatopb.RegisterTomatoServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
