package db

type Order struct {
	ID      string         `gorm:"primaryKey"`
	UserID  string         `gorm:"not null"`
	ItemIDs pq.StringArray `gorm:"type:text[]"`
	Quantities pq.Int64Array `gorm:"type:integer"`
	Status string `gorm:"default:'pending'"`
	DeliveryAddress s
}
