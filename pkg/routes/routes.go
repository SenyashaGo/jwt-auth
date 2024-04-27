package routes

import (
	"net/http"

	"github.com/SenyashaGo/jwt-auth/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Setup(storage *storage.Storage, app *gin.Engine) {
	api := app.Group("/user")
	api.GET("/get-user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello!"})
	})
	api.POST("/register", func(ctx *gin.Context) { Register(ctx, storage) })
	api.POST("/login", func(ctx *gin.Context) {})
	api.POST("/logout", func(ctx *gin.Context) {})
}
