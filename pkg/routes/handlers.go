package routes

import (
	"log"
	"net/http"

	"github.com/SenyashaGo/jwt-auth/pkg/models"
	"github.com/SenyashaGo/jwt-auth/pkg/storage"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context, storage *storage.Storage) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:        data["name"],
		PhoneNumber: data["phone number"],
		Email:       data["email"],
		Password:    password,
	}

	raw, err := storage.RegisterUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, raw)
}

func Login(c *gin.Context, storage *storage.Storage) {
	var data map[string]string

	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
	}

	user := models.User{
		Email:    data["email"],
		Password: []byte(data["password"]),
	}

	raw, err := storage.LoginUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)

	}

	if raw.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"msg": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword(raw.Password, []byte(data["password"])); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "wrong password"})
		return
	}

	c.JSON(http.StatusOK, raw)
}
