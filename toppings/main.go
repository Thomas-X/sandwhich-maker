package main

import (
	butcher "github.com/Thomas-X/sandwhich-maker/butcher/butcherpb"
	"github.com/Thomas-X/sandwhich-maker/lettuce-farmer/lettucefarmerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"github.com/Thomas-X/sandwhich-maker/tomato/tomatopb"
	"github.com/Thomas-X/sandwhich-maker/toppings/toppingspb"
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

func setupButcherConn() (*butcher.ButcherServiceClient, **grpc.ClientConn) {
	butcherPortConn := strconv.Itoa(shared.BUTCHER_PORT)
	log.Printf("Connecting to.. %v", "localhost:"+butcherPortConn)
	conn, err := grpc.Dial("localhost:"+butcherPortConn, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to butcher server %v", err)
	}
	c := butcher.NewButcherServiceClient(conn)
	log.Printf("Created butcher client %v", c)
	return &c, &conn
}

func setupTomatoConn() (*tomatopb.TomatoServiceClient, **grpc.ClientConn) {
	tomatoPortConn := strconv.Itoa(shared.TOMATO_PORT)
	conn, err := grpc.Dial("localhost:"+tomatoPortConn, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to butcher server %v", err)
	}
	c := tomatopb.NewTomatoServiceClient(conn)
	log.Printf("Created tomato client %v", c)
	return &c, &conn
}

func setupLettuceConn() (*lettucefarmerpb.LettuceServiceClient, **grpc.ClientConn) {
	lettucePortConn := strconv.Itoa(shared.LETTUCE_PORT)
	conn, err := grpc.Dial("localhost:"+lettucePortConn, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to butcher server %v", err)
	}
	c := lettucefarmerpb.NewLettuceServiceClient(conn)
	log.Printf("Created tomato client %v", c)
	return &c, &conn
}

func (s *server) ProvideTopping(ctx context.Context, in *toppingspb.ProvideToppingRequest) (*toppingspb.ProvideToppingResponse, error) {
	var meats []*toppingspb.Meat
	var tomatoes []*toppingspb.Tomato
	var lettuces []*toppingspb.Lettuce
	c, conn := setupButcherConn()
	lc, lconn := setupLettuceConn()
	tc, tconn := setupTomatoConn()
	defer (*conn).Close()
	defer (*lconn).Close()
	defer (*tconn).Close()

	// Tomatoes
	bres, berr := (*tc).GetTomato(context.Background(), &tomatopb.GetTomatoRequest{
		Tomatoes: in.Tomatoes,
	})
	if berr != nil {
		log.Fatalf("Couldn't get tomatoes: %v", berr)
	}

	for a := 0; a < len(bres.Tomatoes); a++ {
		tomatoes = append(tomatoes, &toppingspb.Tomato{
			Name:  bres.Tomatoes[a].Name,
			Price: bres.Tomatoes[a].Price,
		})
	}

	// Lettuces
	lres, lerr := (*lc).GetLettuce(context.Background(), &lettucefarmerpb.GetLettuceRequest{
		Lettuces: in.Lettuces,
	})
	if lerr != nil {
		log.Fatalf("Couldn't get lettuces: %v", berr)
	}
	for a := 0; a < len(lres.Lettuces); a++ {
		lettuces = append(lettuces, &toppingspb.Lettuce{
			Name:  lres.Lettuces[a].Name,
			Price: lres.Lettuces[a].Price,
		})
	}

	// Meats
	for i := 0; i < len(in.Meats); i++ {
		bres, berr := (*c).GetMeat(context.Background(), &butcher.GetMeatRequest{
			Name: in.Meats[i],
		})
		if berr != nil {
			log.Fatalf("Couldn't get meats: %v", berr)
		}
		meat := &toppingspb.Meat{
			Price: bres.Price,
			Name:  bres.Name,
		}
		meats = append(meats, meat)
	}

	return &toppingspb.ProvideToppingResponse{
		Meats:    meats,
		Lettuces: lettuces,
		Tomatoes: tomatoes,
	}, nil
}

// TODO lettuce / tomato calls
func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(shared.TOPPING_PORT))
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)
	toppingspb.RegisterToppingServiceServer(s, &server{})
	log.Printf("Serving..")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
