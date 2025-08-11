package main

import (
	"log"
	"net"

	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"
	//Create DB connection
	userRepo, err := db.NewUserRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	//Create gRPC Sever
	lis, err := net.Listen("tcp", ":")
}