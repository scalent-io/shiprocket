package app

import (
	"fmt"
	"log"

	"github.com/scalent-io/shiprocket/apimodel"

	"github.com/spf13/viper"
)

type ShiprocketConfig struct {
	Config apimodel.Config
}

func InitConfig() (*ShiprocketConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./app")
	viper.SetEnvPrefix("shiprocket")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var config ShiprocketConfig

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Println("Unmarshal-----------", err)
		log.Fatal(err)
		return nil, err
	}

	return &config, err
}
