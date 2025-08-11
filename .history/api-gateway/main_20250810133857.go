package main

import (
	"log"
	userpb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := proto.
}