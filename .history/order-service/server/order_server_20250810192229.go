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
		return nil, fmt.Errorf("at least one item is required")
	}

	// Set default quantities if not provided
	quantities := req.Quantities
	if len(quantities) == 0 {
		quantities = make([]int32, len(req.ItemIds))
		for i := range quantities {
			quantities[i] = 1 //default quantity
		}
		if len(quantities) != len(req.ItemIds) {
		return nil, fmt.Errorf("quantities length must match items length")
	}
	// Calculate total amount (you might want to fetch actual prices from menu service)
	
}