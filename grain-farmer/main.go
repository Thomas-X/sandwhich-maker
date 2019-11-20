package main

import (
	"github.com/Thomas-X/sandwhich-maker/fertilizer/fertilizerpb"
	"github.com/Thomas-X/sandwhich-maker/grain-farmer/grainfarmerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
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

func (s *server) GetGrain(ctx context.Context, in *grainfarmerpb.GetGrainRequest) (*grainfarmerpb.GetGrainResponse, error) {
	grain := 0
	price := 0
	switch in.BreadName {
	case "Soleil":
		grain = 200
		price = 5
	default:
		grain = 100
		price = 2
	}

	c, conn := setupFertilizer()
	defer (*conn).Close()

	res, err := (*c).GetFertilizer(context.Background(), &fertilizerpb.GetFertilizerRequest{
		Amount: int32(price),
	})
	if err != nil {
		log.Fatalf("Couldn't get fertilizer for grain %v", err)
	}
	price = int(res.Price)

	return &grainfarmerpb.GetGrainResponse{
		Grain: int32(grain),
		Price: int32(price),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.GRAIN_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	grainfarmerpb.RegisterGrainFarmerServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
