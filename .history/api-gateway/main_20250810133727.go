package main

import (
	"log"

	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal
	}
}