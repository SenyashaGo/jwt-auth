package storage

import (
	"github.com/SenyashaGo/jwt-auth/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage(dsn string) (*Storage, error) {
	connect, err := gorm.Open(postgres.Open(dsn), nil)
	if err != nil {
		return nil, err
	}

	err = connect.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return &Storage{connect}, nil
}

func (s *Storage) RegisterUser(user models.User) (models.User, error) {
	tx := s.db.Create(&user)
	return user, tx.Error
}

func (s *Storage) LoginUser(user models.User) (models.User, error) {
	tx := s.db.Where("email = ?", user.Email).First(&user)
	return user, tx.Error
}
