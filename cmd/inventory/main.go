package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"cyaniccerulean.com/inventory/v2/internal/model"
	"cyaniccerulean.com/inventory/v2/internal/service"
	"github.com/spf13/viper"
)

// bootstrap application
func main() {
	// set up signal trap
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// read config from environment
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// create new service
	service, err := service.New(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	// start serving requests
	if err = service.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

// read config from environment and config file
// environment overrides values from file
func loadConfig() (config model.Config, err error) {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
