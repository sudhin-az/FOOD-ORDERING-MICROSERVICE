package db

import "time"

type Order struct {
	ID              string         `gorm:"primaryKey"`
	UserID          string         `gorm:"not null"`
	ItemIDs         pq.StringArray `gorm:"type:text[]"`
	Quantities      pq.Int64Array  `gorm:"type:integer"`
	Status          string         `gorm:"default:'pending'"`
	DeliveryAddress string         `gorm:"type:text"`
	TotalAmount     float64        `gorm:"type:decimal(10, 2)"`
	CreateAt        time.Time `gorm:"autoCreateTime"`
}
