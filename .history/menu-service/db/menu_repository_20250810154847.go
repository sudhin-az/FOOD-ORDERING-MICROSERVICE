package db

import (
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
	if err := 
}
