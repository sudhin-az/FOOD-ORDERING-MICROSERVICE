package server

import (
	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
	"gorm.io/gorm"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo db) *UserServer {
	return &UserServer{repo: }
}