package api

import (
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
	"vps-provider/config"
)

var log = logging.Logger("api")

func ConfigRouter(router *gin.Engine, cfg config.Config) {
	apiV1 := router.Group("/api/v1")
	authMiddleware, err := jwtGinMiddleware(cfg.SecretKey)
	if err != nil {
		log.Fatalf("jwt auth middleware: %v", err)
	}

	err = authMiddleware.MiddlewareInit()
	if err != nil {
		log.Fatalf("authMiddleware.MiddlewareInit: %v", err)
	}

	storage := apiV1.Group("/storage")
	storage.POST("/get_verify_code", GetVerifyCodeHandle)
	storage.POST("/register", UserRegister)
	storage.POST("/password_reset", PasswordRest)
	storage.POST("/login", authMiddleware.LoginHandler)
	storage.POST("/logout", authMiddleware.LogoutHandler)
	storage.Use(authMiddleware.MiddlewareFunc())
	storage.GET("/refresh_token", authMiddleware.RefreshHandler)
}