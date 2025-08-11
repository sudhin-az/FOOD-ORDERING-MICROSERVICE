package main

import (
	"log"
	userpb "github.com/sudhin-az/FOOD-ORDERING/user-service/pro
	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := proto.
}