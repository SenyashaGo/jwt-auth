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
