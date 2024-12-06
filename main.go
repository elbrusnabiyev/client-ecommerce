package main

import (
	"context"
	"log"
	"time"

	pb "github.com/elbrusnabiyev/client-ecommerce/ecommerce"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addres = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(addres, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	desscription := `Meet Apple iPhone 11. All-new dual-camera system with Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: desscription, Price: price}) // pos.5
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added succesfully", r.Value)

	product, err := c.GetProduct(ctx, &pb.ProductID{Value: r.Value}) // pos.6
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: ", product.String())

}
