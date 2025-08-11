package main

import (
	"log"
	"net"

	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
	"github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"github.com/sudhin-az/FOOD-ORDERING/user-service/server"
	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"
	//Create DB connection
	userRepo, err := db.NewUserRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	//Create gRPC Sever
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	//create gRPC server and register service
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server.NewUserServer(userRepo))
}