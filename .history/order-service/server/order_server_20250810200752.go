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
	if req.UserId == "" {
		return nil, fmt.Errorf("user_id is required")
	}

	orders, err := s.repo.GetOrderByUser(ctx, req.UserId)
	if err != nil {
		return  nil, fmt.Errorf("failed to get orders for user %v", err)
	}
	var pbOrders []*pb.Order
	for _, o := range orders {
		quantities := make([]int32, len(o.Quantities))
		for i, q := range o.Quantities {
			quantities[i] = int32(q)
		}
		pbOrders = append(pbOrders, &pb.Order{
			OrderId:         o.ID,
			UserId:          o.UserID,
			ItemIds:         o.ItemIDs,
			Quantities:      quantities,
			Status:          o.Status,
			DeliveryAddress: o.DeliveryAddress,
			TotalAmount:     o.TotalAmount,
			CreatedAt:       o.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:       o.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	} 
	return &pb.OrderList{Orders: pbOrders,
		TotalCount: int32(len(pbOrders)),
		}, nil
}

func (s *OrderServer) GetOrderByID(ctx context.Context, req *pb.OrderRequest) (*pb.OrderResponse, error) {
	if req.OrderId == "" {
		return nil, fmt.Errorf("order_id is required")
	}
	order, err := s.repo.GetOrderByID(ctx, req.OrderId)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve order: %v", err)
	}
	quantities := make([]int32, len(order.Quantities))
	for i, q := range order.Quantities {
		quantities[i] = int32(q)
	}
	return &pb.OrderResponse{
		OrderId:         order.ID,
		UserId:          order.UserID,
		ItemIds:         order.ItemIDs,
		Quantities:      quantities,
		Status:          order.Status,
		DeliveryAddress: order.DeliveryAddress,
		TotalAmount:     order.TotalAmount,
		CreatedAt:       order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:       order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}

func (s *OrderServer) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderStatusRequest) (*pb.UpdateOrderStatusResponse, error) {
	if req.OrderId == "" {
		return nil, fmt.Errorf("order_id is required")
	}
	if req.Status == "" {
		return nil, fmt.Errorf("status is required")
	}

	//Validate status 
	validateStatuses := map[string]bool{
		"pending": true,
		"confirmed": true,
		"delivered": true,
		"cancelled": true,
	}
	if !validateStatuses[req.Status] {
		return nil, fmt.Errorf("invalid status: %s", req.Status)
	}
	err := s.repo.UpdateOrderStatus(ctx, req.OrderId, req.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update order status: %v", err)
	}
	return &pb.UpdateOrderStatusResponse{
		OrderId: req.OrderId,
		Status: req.Status,
		Message: "Order status updated successfully",
	}, nil
}

func (s *OrderServer) GetAllOrders() {
	
}