package server

import (
	"context"
	"fmt"

	"github.com/sudhin-az/FOOD-ORDERING/order-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/order-service/proto"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	repo *db.OrderRepository
}

func (s *OrderServer) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	
	if req.UserId == "" {
		return nil, fmt.Errorf("user_id is required")
	}
	if len(req.ItemIds) == 0 {
		return nil, fmt.Errorf("at least one item is re")
	}
}