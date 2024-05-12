package routes

import (
	"github.com/SenyashaGo/jwt-auth/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Setup(storage *storage.Storage, app *gin.Engine) {
	api := app.Group("/user")
	api.GET("/get-user", func(ctx *gin.Context) { User(ctx, storage) })
	api.POST("/register", func(ctx *gin.Context) { Register(ctx, storage) })
	api.POST("/login", func(ctx *gin.Context) { Login(ctx, storage) })
	api.POST("/logout", func(ctx *gin.Context) { Logout(ctx) })
}
