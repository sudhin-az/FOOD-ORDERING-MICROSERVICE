package server

import (
	"context"

	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context) {
	
}