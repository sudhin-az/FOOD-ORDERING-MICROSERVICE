package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	menupb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"
	userpb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"google.golang.org/grpc"
)

func main() {
	// create grpc connections
	userConn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	userClient := userpb.NewUserServiceClient(userConn)

	menuConn, _ := grpc.Dial("localhost:50052", grpc.WithInsecure())
	menuClient := menupb.NewMenuServiceClient(menuConn)

	orderConn, _ := grpc.Dial("localhost:50053", grpc.withi)

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

	r.POST("/addItem", func(ctx *gin.Context) {
		var req menupb.AddItemRequest
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res, _ := menuClient.AddItem(context.Background(), &req)
		ctx.JSON(http.StatusOK, res)
	})
	r.GET("/getAllItems", func(ctx *gin.Context) {
		menus, err := menuClient.GetAllItems(ctx, &menupb.Empty{})
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		ctx.JSON(http.StatusOK, menus)
	})

	log.Println("API Gateway is running on :8080")
	r.Run(":8080")
}
