package main

import (
	"context"
	"github.com/Thomas-X/sandwhich-maker/baker/bakerpb"
	"github.com/Thomas-X/sandwhich-maker/shared"
	"github.com/Thomas-X/sandwhich-maker/toppings/toppingspb"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func setupToppingsClient() (*toppingspb.ToppingServiceClient, **grpc.ClientConn) {
	toppingsPort := strconv.Itoa(shared.TOPPING_PORT)
	conn, err := grpc.Dial("localhost:"+toppingsPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to toppings server %v", err)
	}
	c := toppingspb.NewToppingServiceClient(conn)
	log.Printf("Created toppings client %v", c)
	return &c, &conn
}

func setupBreadClient() (*bakerpb.BakerServiceClient, **grpc.ClientConn) {
	breadPort := strconv.Itoa(shared.BAKER_PORT)
	conn, err := grpc.Dial("localhost:"+breadPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't start listening to toppings server %v", err)
	}
	c := bakerpb.NewBakerServiceClient(conn)
	log.Printf("Created toppings client %v", c)
	return &c, &conn
}

func main() {
	t, tconn := setupToppingsClient()
	b, bconn := setupBreadClient()
	defer (*tconn).Close()
	defer (*bconn).Close()
	log.Printf("ButcherPort: %v", strconv.Itoa(shared.BUTCHER_PORT))

	bres, err1 := (*b).GetBread(context.Background(), &bakerpb.GetBreadRequest{
		Name:   "Soleil",
		Amount: 2,
	})
	tres, err2 := (*t).ProvideTopping(context.Background(), &toppingspb.ProvideToppingRequest{
		Meats: []string{
			"roastbeef",
			"salami",
		},
		Lettuces: []string{
			"iceberg",
			"kale",
		},
		Tomatoes: []string{
			"dutch",
			"roma",
		},
	})

	if err1 != nil {
		log.Fatalf("Error getting bread: %v", err1)
	}
	if err2 != nil {
		log.Fatalf("Error getting toppings: %v", err2)
	}
	log.Printf("Got tres %v", tres)
	log.Printf("Got bres %v", bres)
}
