package main

import (
	"log"

	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
)

func main() {
	dbURL := "postgres://postgres:sudhin123@localhost:5432/food_ordering?sslmode=disable"
	userRepo, err := db.NewUserRepository(dbURL)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
}