package main

import (
	"log"
	"net"

	"github.com/sudhin-az/E-LEARNING-GRPC/user-service/server"
	"github.com/sudhin-az/FOOD-ORDERING/menu-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"
	"google.golang.org/grpc"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"
	//Create DB connection
	repo, err := db.NewMenuRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	//Create gRPC Sever
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	//create gRPC server and register service
	s := grpc.NewServer()
	pb(s, server.NewUserServer(repo))

	log.Println("UserService running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
