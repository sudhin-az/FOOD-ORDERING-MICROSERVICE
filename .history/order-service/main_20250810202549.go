package main

import (
	"log"
	"net"

	"github.com/sudhin-az/FOOD-ORDERING/order-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/order-service/proto"
	"github.com/sudhin-az/FOOD-ORDERING/order-service/server"
	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"

	//Create DB Connection
	repo, err := db.NewOrderRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	//Create gRpc Server
	lis, err := net.Listen("tcp", "50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderServiceServer(s, server.)
}
