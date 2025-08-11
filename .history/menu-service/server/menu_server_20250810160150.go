package server

import pb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"

type MenuServer struct {
	pb.UnimplementedMenuServiceServer
	repo *d
}