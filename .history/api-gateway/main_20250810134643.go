package main

import (
	"context"
	"net/http"

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
		if err := c(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := userClient.CreateUser(context.Background(), &req)
		c.JSON(http.StatusOK, res)
	})

	r.GET("/getusers", func(ctx *gin.Context) {
		users, err := userClient.GetAllUsers(ctx, &userpb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, users)

	})
}
