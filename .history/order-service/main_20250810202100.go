package main

import (
	"net"

	"github.com/sudhin-az/FOOD-ORDERING/order-service/db"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"

	//Create DB Connection
	repo, err := db.NewOrderRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}

	//Create gRpc Server
	lis, err := net.Listen("")
}