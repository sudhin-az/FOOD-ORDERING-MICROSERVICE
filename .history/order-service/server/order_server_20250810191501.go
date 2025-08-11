package server

import pb "github.com/sudhin-az/FOOD-ORDERING/order-service/proto"

type OrderServer struct {
	pb.UnimplementedOrderServiceServer
	repo 
}