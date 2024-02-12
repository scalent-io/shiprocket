package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	// token, err := auth.GetToken(config.Config.BaseURL, config.Config.Email, config.Config.Password)
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("err", err)
	// }

	fmt.Println("config :-----", config)
}

func GetShiprocketClient(ctx context.Context, shiprocketConfig ShiprocketConfig) (*ShiprocketClient, error) {

	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	ShiprocketClient, err := initServer()
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	ShiprocketClient.Options.Token, err = ShiprocketClient.Options.ShiprocketService.GetToken(config.Config.BaseURL, config.Config.Email, config.Config.Password)
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	return ShiprocketClient, nil
}
