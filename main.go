package main

import (
	"context"
	"fmt"
	logging "github.com/ipfs/go-log/v2"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vps-provider/api"
	"vps-provider/config"
	"vps-provider/core/dao"
	"vps-provider/core/oplog"
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

	if err := dao.Init(&cfg); err != nil {
		log.Fatalf("initital: %v\n", err)
	}

	oplog.Subscribe(context.Background())

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
