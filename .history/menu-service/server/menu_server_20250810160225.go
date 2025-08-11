package server

import (
	"github.com/sudhin-az/FOOD-ORDERING/menu-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/menu-service/proto"
)

type MenuServer struct {
	pb.UnimplementedMenuServiceServer
	repo *db.MenuRepository
}

func NewMenuServer(repo *db.MenuRepository) *me {
	
}