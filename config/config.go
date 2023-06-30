package config

// Cfg config
var Cfg Config

// Config server config
type Config struct {
	Mode          string
	APIListen     string
	DatabaseURL   string
	SecretKey     string
	RedisAddr     string
	RedisPassword string
	Email         EmailConfig
}

// EmailConfig email
type EmailConfig struct {
	Name     string
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}
