package db

import (
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Order struct {
	ID              string         `gorm:"primaryKey"`
	UserID          string         `gorm:"not null"`
	ItemIDs         pq.StringArray `gorm:"type:text[]"`
	Quantities      pq.Int64Array  `gorm:"type:integer"`
	Status          string         `gorm:"default:'pending'"`
	DeliveryAddress string         `gorm:"type:text"`
	TotalAmount     float64        `gorm:"type:decimal(10, 2)"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime"`
}

type OrderRepository struct {
	conn *gorm.DB
}

func NewOrderRepository(dbURL string) (*OrderRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}
	if err := conn.AutoMigrate(&Order{}); err != nil {
		return nil, fmt.Errorf("failed to migrate order table: %v", err)
	}
	return o
}
