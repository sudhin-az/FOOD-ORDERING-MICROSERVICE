package main

import "github.com/sudhin-az/FOOD-ORDERING/order-service/db"

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"

	//Create DB Connection
	repo, err := db.NewOrderRepository(dbURL)
}