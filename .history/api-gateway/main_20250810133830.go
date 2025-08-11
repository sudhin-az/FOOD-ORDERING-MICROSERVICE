package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := pro
}