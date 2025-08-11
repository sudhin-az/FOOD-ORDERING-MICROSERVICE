package server

import (
	"context"
	"fmt"

	"github.com/sudhin-az/FOOD-ORDERING/user-service/db"
	pb "github.com/sudhin-az/FOOD-ORDERING/user-service/proto"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
	repo *db.UserRepository
}

func NewUserServer(repo *db.UserRepository) *UserServer {
	return &UserServer{repo: repo}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.repo.CreateUser(ctx, req.Name, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create User %v", err)
	}
	return &pb.CreateUserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserList, error) {
	users, err := s.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve Users: %v", err)
	}
	var pbUsers *[]pb.User
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id: u.ID,
			Name: u.Name,
			Email: u.Email,
		})
	}
	return &pb.UserList{Users: pbUsers}, n
}