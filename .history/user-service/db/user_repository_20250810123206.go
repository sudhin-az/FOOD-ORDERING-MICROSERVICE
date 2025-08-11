package db

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

type UserRepository struct


