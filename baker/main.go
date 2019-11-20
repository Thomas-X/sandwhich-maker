package main

import (
	"github.com/Thomas-X/sandwhich-maker/baker/bakerpb"
	"github.com/Thomas-X/sandwhich-maker/grain-farmer/grainfarmerpb"
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

func setupGrainFarmerConn() (*grainfarmerpb.GrainFarmerServiceClient, **grpc.ClientConn) {
	grainFarmerPort := strconv.Itoa(shared.GRAIN_PORT)
	conn, err := grpc.Dial("localhost:"+grainFarmerPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to grain server %v", err)
	}
	c := grainfarmerpb.NewGrainFarmerServiceClient(conn)
	log.Printf("Created grain farmer client %v", c)
	return &c, &conn
}

func setupMoonSeedFarmerConn() (*moonseedfarmerpb.MoonSeedFarmerServiceClient, **grpc.ClientConn) {
	moonSeedFarmerPort := strconv.Itoa(shared.MOONSEED_PORT)
	conn, err := grpc.Dial("localhost:"+moonSeedFarmerPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to moonseed server %v", err)
	}
	c := moonseedfarmerpb.NewMoonSeedFarmerServiceClient(conn)
	log.Printf("Created moonseed farmer client %v", c)
	return &c, &conn
}

func (s *server) GetBread(ctx context.Context, in *bakerpb.GetBreadRequest) (*bakerpb.GetBreadResponse, error) {
	name := ""
	price := int32(0)
	switch in.Name {
	case "Soleil":
		name = "Soleil"
		m, mconn := setupMoonSeedFarmerConn();
		defer (*mconn).Close()
		moonseedres, err := (*m).GetMoonSeed(context.Background(), &moonseedfarmerpb.GetMoonSeedRequest{BreadName:name})
		if err != nil {
			log.Fatalf("Error getting moonseed %v", err)
		}
		if moonseedres.Price != 0 {
			price += moonseedres.Price
		}
	}
	c, conn := setupGrainFarmerConn();
	defer (*conn).Close()

	grainres, err := (*c).GetGrain(context.Background(), &grainfarmerpb.GetGrainRequest{
		BreadName: name,
	})
	if err != nil {
		log.Fatalf("Error getting grain: %v", err)
	}
	if grainres.Price != 0 {
		price += grainres.Price
	}

	return &bakerpb.GetBreadResponse{
		Bread: &bakerpb.Bread{
			Name:   name,
			Amount: in.Amount,
			Price:  price,
			GrainUsed: grainres.Grain,
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:" + strconv.Itoa(shared.BAKER_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}
	s := grpc.NewServer()
	bakerpb.RegisterBakerServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
