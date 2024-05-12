package storage

import (
	"github.com/SenyashaGo/jwt-auth/pkg/models"
	"github.com/golang-jwt/jwt"
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

func (s *Storage) GetUser(claims *jwt.StandardClaims) (models.User, error) {
	var user models.User
	tx := s.db.Where("id = ?", claims.Issuer).First(&user)
	return user, tx.Error
}
