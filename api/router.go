package api

import (
	"time"

	"vps-provider/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("api")

type Server struct {
	cfg    config.Config
	router *gin.Engine
}

// NewServer new a router server
func NewServer(cfg config.Config) (*Server, error) {
	gin.SetMode(cfg.Mode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowAllOrigins:  true,
	}))

	configRouter(router, cfg)

	s := &Server{
		cfg:    cfg,
		router: router,
	}

	return s, nil
}

// Run server
func (s *Server) Run() {
	err := s.router.Run(s.cfg.APIListen)
	if err != nil {
		log.Fatal(err)
	}
}

// Close server
func (s *Server) Close() {
}

func configRouter(router *gin.Engine, cfg config.Config) {
	apiV1 := router.Group("/api/v1")
	// authMiddleware, err := jwtGinMiddleware(cfg.SecretKey)
	// if err != nil {
	// 	log.Fatalf("jwt auth middleware: %v", err)
	// }

	// err = authMiddleware.MiddlewareInit()
	// if err != nil {
	// 	log.Fatalf("authMiddleware.MiddlewareInit: %v", err)
	// }

	apiV1.GET("/", homePage)
	// apiV1.POST("/register", userRegisterHandler)
	// apiV1.POST("/login", authMiddleware.LoginHandler)
	// apiV1.POST("/logout", authMiddleware.LogoutHandler)
	apiV1.GET("/describe_price", describePrice)
	apiV1.GET("/describe_available_resource", describeAvailableResource)
	apiV1.GET("/describe_instance_type", describeRecommendInstanceType)
	apiV1.GET("/describe_images", describeImages)
	apiV1.GET("/describe_regions", describeRegions)
	// apiV1.GET("/create_security_group", createSecurityGroup)
	apiV1.GET("/create_instance", createInstance)
	// apiV1.Use(authMiddleware.MiddlewareFunc())
	// apiV1.GET("/refresh_token", authMiddleware.RefreshHandler)
}
