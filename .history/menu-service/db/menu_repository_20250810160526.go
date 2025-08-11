package db

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuItem struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price float64
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

func (r *MenuRepository) AddItem(ctx context.Context, name string, price f) (MenuItem, error) {
	id := uuid.NewString()

	newMenu := MenuItem{
		ID:    id,
		Name:  name,
		Price: price,
	}
	err := r.conn.Create(&newMenu).Error
	if err != nil {
		return MenuItem{}, fmt.Errorf("failed to Add Item: %v", err)
	}
	log.Printf("Item Added Successfully: %+v", newMenu)
	return newMenu, nil
}

func (r *MenuRepository) GetAllItems(ctx context.Context) ([]MenuItem, error) {
	var menu []MenuItem

	err := r.conn.Find(&menu).Error
	if err != nil {
		return []MenuItem{}, fmt.Errorf("failed to get menuItems: %+v", err)
	}
	log.Printf("Items Retrieved Successfully: %+v", menu)
	return menu, nil
}
