package main

import (
	"context"
	"fmt"
	"log"

	pb "test-grpc-client/protos/product"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductServiceClient(conn)

	req := pb.ProductNameRequest{
		Id: 2,
	}

	// call Greet service
	res, err := c.SendProductName(context.Background(), &req)
	if err != nil {
		log.Fatalf("request failed: %v", err)
	}
	log.Println("client get product name")
	fmt.Println(res)

}
