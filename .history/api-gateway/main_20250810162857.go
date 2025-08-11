package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	userpb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	menupb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := userpb.NewUserServiceClient(userConn)

	menuConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	menuClient := menupb.NewMenuServiceClient(menuConn)

	r := gin.Default()

	r.POST("/createuser", func(ctx *gin.Context) {
		var req userpb.CreateUserRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := userClient.CreateUser(context.Background(), &req)
		ctx.JSON(http.StatusOK, res)
	})

	r.GET("/getusers", func(ctx *gin.Context) {
		users, err := userClient.GetAllUsers(ctx, &userpb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, users)

	})

	r.POST("/addItem")

	log.Println("API Gateway is running on :8080")
	r.Run(":8080")
}
