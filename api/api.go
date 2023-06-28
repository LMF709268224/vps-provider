package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"vps-provider/config"
	"vps-provider/utils"
)

type Server struct {
	cfg    config.Config
	router *gin.Engine
}

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
	ConfigRouter(router, cfg)
	s := &Server{
		cfg:    cfg,
		router: router,
	}

	return s, nil
}

func (s *Server) Run() {
	err := s.router.Run(s.cfg.ApiListen)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Close() {
}

func (s *Server) sendEmail(sendTo string, registrations []string) error {
	var EData utils.EmailData
	EData.Subject = "[Application]: Your Device Info"
	EData.Tittle = "please check your device id "
	EData.SendTo = sendTo
	EData.Content = "<h1>Your Device ID ：</h1>\n"
	for _, registration := range registrations {
		EData.Content += registration + "<br>"
	}

	err := utils.SendEmail(s.cfg.Email, EData)
	if err != nil {
		return err
	}
	return nil
}
