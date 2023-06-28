package config

var Cfg Config

type Config struct {
	Mode          string
	ApiListen     string
	DatabaseURL   string
	SecretKey     string
	RedisAddr     string
	RedisPassword string
	Email         EmailConfig
}

type EmailConfig struct {
	Name     string
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}
