package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"vps-provider/api"
	"vps-provider/config"
	"vps-provider/storage/mysql"

	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/viper"
)

func main() {
	OsSignal := make(chan os.Signal, 1)

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("reading config file: %v\n", err)
	}

	var cfg config.Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unmarshaling config file: %v\n", err)
	}
	config.Cfg = cfg
	if cfg.Mode == "debug" {
		logging.SetDebugLogging()
	}

	if err := mysql.Init(&cfg); err != nil {
		log.Fatalf("initital: %v\n", err)
	}

	srv, err := api.NewServer(cfg)
	if err != nil {
		log.Fatalf("create api server: %v\n", err)
	}
	go srv.Run()

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM)
	_ = <-OsSignal
	srv.Close()

	fmt.Printf("Exiting received OsSignal\n")
}
