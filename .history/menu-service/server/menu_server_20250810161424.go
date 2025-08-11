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
		return nil, fmt.Errorf("failed to Add Item %v", err)
	}
	return &pb.AddItemResponse{
		Id: menu.ID,
		Name: menu.Name,
		Price: menu.Price,
	}, nil
}

func (s *MenuServer) GetAllItems(ctx context.Context, req *pb.Empty) (*pb.ItemList, error) {
	menu, err := s.repo.GetAllItems(ctx)
	if err != nil {
		fmt.Errorf("failed to get Items %v", err)
	}
	var pbMenu []*pb.Item
	for _, m := range menu {
		pbMenu = append(pbMenu, &pb.Item{
			i: m.ID,
			Name: m.Name,
			Price: m.Price,
		})
	}
	return &pb.ItemList{Items: pbMenu}, nil
}