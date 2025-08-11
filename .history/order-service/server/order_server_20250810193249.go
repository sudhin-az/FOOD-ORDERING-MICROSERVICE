package server

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sudhin-az/FOOD-ORDERING/order-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/order-service/proto"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	repo *db.OrderRepository
}

func (s *OrderServer) PlaceOrder(ctx context.Context, req *pb.PlaceOrderRequest) (*pb.PlaceOrderResponse, error) {
	// Validate input
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
			quantities[i] = 1 // default quantity
		}
	}

	// Validate quantities length matches items length
	if len(quantities) != len(req.ItemIds) {
		return nil, fmt.Errorf("quantities length must match items length")
	}

	// Calculate total amount (you might want to fetch actual prices from menu service)
	totalAmount := 0.0
	for _, qty := range quantities {
		totalAmount += float64(qty) 
	}

	order, err := s.repo.PlaceOrder(ctx, req.UserId, req.ItemIds, quantities, req.DeliveryAddress, totalAmount)
	if err != nil {
		return nil, fmt.Errorf("failed to place order: %v", err)
	}

	return &pb.PlaceOrderResponse{
		OrderId:     order.ID,
		Status:      order.Status,
		Message:     "Order placed successfully",
		TotalAmount: order.TotalAmount,
	}, nil
}

func (s *OrderServer) GetOrderByUser(ctx context.Context, req *pb.UserRequest) (*pb.OrderList, error) {
	userID := uuid.NewString()

	orders, err := s.repo.GetOrderByUser(ctx, req.UserId)
}