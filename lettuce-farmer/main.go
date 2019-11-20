package main

import (
	"github.com/Thomas-X/sandwhich-maker/fertilizer/fertilizerpb"
	"github.com/Thomas-X/sandwhich-maker/lettuce-farmer/lettucefarmerpb"
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

func GetFertilizerForLettuce(amount int32, name string, c *fertilizerpb.FertilizerServiceClient) *lettucefarmerpb.Lettuce {
	lettuce, err := (*c).GetFertilizer(context.Background(), &fertilizerpb.GetFertilizerRequest{
		Amount: amount,
	})
	if err != nil {
		log.Fatalf("Couldn't get fertilizer for lettuce: %v", err)
	}

	return &lettucefarmerpb.Lettuce{
		Name:  name,
		Price: lettuce.Price,
	}
}

func (s *server) GetLettuce(ctx context.Context, in *lettucefarmerpb.GetLettuceRequest) (*lettucefarmerpb.GetLettuceResponse, error) {
	var lettuces []*lettucefarmerpb.Lettuce
	c, conn := setupFertilizer()
	defer (*conn).Close()

	for i := 0; i < len(in.Lettuces); i++ {
		switch in.Lettuces[i] {
		case "iceberg":
			lettuces = append(lettuces, GetFertilizerForLettuce(4, "iceberg", c))
		case "romaine":
			lettuces = append(lettuces, GetFertilizerForLettuce(8, "romaine", c))
		case "kale":
			lettuces = append(lettuces, GetFertilizerForLettuce(2, "kale", c))
		}
	}
	return &lettucefarmerpb.GetLettuceResponse{Lettuces: lettuces}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.LETTUCE_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	lettucefarmerpb.RegisterLettuceServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
