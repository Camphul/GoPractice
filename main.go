package main

import (
	"github.com/camphul/gopractice/config"
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("Hello GoPractice")
	loadViperConfig()
}

//Check how to use config files(using viper libary)
func loadViperConfig() {
	viper.SetConfigName("viper-config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("General Config Name: %s", configuration.General.Name)
	log.Printf("General Config Version: %s", configuration.General.Version)
}
