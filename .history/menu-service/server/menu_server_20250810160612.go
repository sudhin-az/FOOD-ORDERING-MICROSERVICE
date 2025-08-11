package server

import (
	"context"
	"fmt"

	"github.com/sudhin-az/FOOD-ORDERING/menu-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"
)

type MenuServer struct {
	pb.UnimplementedMenuServiceServer
	repo *db.MenuRepository
}

func NewMenuServer(repo *db.MenuRepository) *MenuServer {
	return &MenuServer{repo: repo}
}

func (s *MenuServer) AddItem(ctx context.Context, req *pb.AddItemRequest) (*pb.AddItemResponse, error) {
	menu, err := s.repo.AddItem(ctx, req.Name, req.Price)
	if err != nil {
		return nil, fmt.e
	}
}