package main

import (
	"log"

	"github.com/gin-gonic/gin"
	userpb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := userpb.NewUserServiceClient(userConn)

	r := gin.Default()

	r.POST("/createuser", func(ctx *gin.Context) {
		var req userpb.CreateUserRequest
	})
}