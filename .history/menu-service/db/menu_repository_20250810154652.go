package db

import "gorm.io/gorm"

type MenuItem struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Price string
}

type MenuRepository struct {
	conn *gorm.DB
}

func NewMenuRepository(dbURL string) (*MenuRepository, error) {
	
}
