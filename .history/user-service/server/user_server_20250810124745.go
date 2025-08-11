package server

import (
	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUser()  {
	
}