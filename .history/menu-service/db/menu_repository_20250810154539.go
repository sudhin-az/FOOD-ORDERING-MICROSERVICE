package db

type MenuItem struct {
	ID string `gorm:"primaryKey"`
	Name string
	Price string
}

type MenuRepository struct {
	conn *g
}