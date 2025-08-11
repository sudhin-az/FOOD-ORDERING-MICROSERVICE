package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price string
}

type MenuRepository struct {
	conn *gorm.DB
}

func NewMenuRepository(dbURL string) (*MenuRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to db: %v", err)
	}
	if err := conn.AutoMigrate(&MenuItem{}); err != nil {
		return nil, fmt.Errorf("failed to migrate menu table: %v", err)
	}
	return &MenuRepository{conn: conn}, nil
}

func (r *MenuRepository) AddItem(ctx context.Context, name, price string) (MenuItem, error) {
	id := uuid.NewString()

	newMenu := MenuItem{
		ID: id,
	}
}
