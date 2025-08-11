package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(dbURL string) (*UserRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %v", err)
	}
	if err := conn.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("failed to migrate user table: %v", err)
	}
	return &UserRepository{conn: conn}, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, name, email string) (User, error) {
	id := uuid.NewString()

	newUser := User{
		ID: id,
		Name: name,
		Email: email,
	}
	err := r.conn.WithContext(ctx).Create(&newUser).Error
	if err != nil {
		return User{}, fmt.Errorf("failed to create user: %v", err)
	}
	log.Printf("User Created Successfully: %+v", newUser)
	return newUser, nil
}


