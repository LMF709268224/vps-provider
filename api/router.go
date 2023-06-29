package api

import (
	"vps-provider/config"

	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("api")

func configRouter(router *gin.Engine, cfg config.Config) {
	apiV1 := router.Group("/api/v1")
	authMiddleware, err := jwtGinMiddleware(cfg.SecretKey)
	if err != nil {
		log.Fatalf("jwt auth middleware: %v", err)
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		log.Fatalf("authMiddleware.MiddlewareInit: %v", err)
	}

	apiV1.POST("/register", userRegisterHandler)
	apiV1.Use(middlewareRole)
	apiV1.POST("/password_reset", resetPasswordHandler)
	apiV1.POST("/login", authMiddleware.LoginHandler)
	apiV1.POST("/logout", authMiddleware.LogoutHandler)
	apiV1.Use(authMiddleware.MiddlewareFunc())
	apiV1.GET("/refresh_token", authMiddleware.RefreshHandler)
}
