package routes

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/SenyashaGo/jwt-auth/pkg/models"
	"github.com/SenyashaGo/jwt-auth/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
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

const SecretKey = "secret"

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

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(raw.ID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "could not login"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, raw)
}

func User(c *gin.Context, storage *storage.Storage) {
	cookie, err := c.Cookie("jwt")

	if err != nil {
		cookie = "Not Set"
		c.JSON(http.StatusBadRequest, cookie)
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	raw, err := storage.GetUser(claims)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, raw)
}

func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, "success")
}
