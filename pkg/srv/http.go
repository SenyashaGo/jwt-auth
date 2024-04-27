package srv

import (
	"net"

	"github.com/SenyashaGo/jwt-auth/pkg/config"
	"github.com/SenyashaGo/jwt-auth/pkg/routes"
	"github.com/SenyashaGo/jwt-auth/pkg/storage"
	"github.com/gin-gonic/gin"
)

func Run(storage *storage.Storage, cfg *config.Config) error {
	app := gin.New()
	routes.Setup(storage, app)
	return app.Run(net.JoinHostPort(cfg.Host, cfg.Port))
}
