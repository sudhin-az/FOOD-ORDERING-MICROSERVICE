package db

type MenuItem struct {
	ID string `gorm:"primaryKey"`
	Name string
}