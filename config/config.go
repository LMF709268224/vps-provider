package config

// Cfg config
var Cfg Config

// Config server config
type Config struct {
	Mode                  string
	APIListen             string
	SecretKey             string
	DryRun                bool
	AliyunAccessKeyID     string
	AliyunAccessKeySecret string
}
