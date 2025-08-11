package server

import (
	"context"

	"github.com/sudhin-az/FOOD-ORDERING/order-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/order-service/proto"
)

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	repo *db.OrderRepository
}

func (s *OrderServer) PlaceOrder(ctx context.Context, p) {
	
}