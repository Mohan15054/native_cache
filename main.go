package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"cachesystem/internal/app"
)

func main() {
	// Load configuration
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		os.Exit(1)
	}

	// Start the cache application
	app.Run()
}
