package db

type User struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(dbURL string) (*UserRepository, error) {
	conn, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %v", err)
	}
	if err := conn.Automigrate(&User{}).Error; err != nil {
		return nil, fmt.Errorf("failed to migrate user table: %v", err)
	}
	return &UserRepository{conn: conn}, nil
}

func (r *)  {
	
}