package main

import (
	"fmt"
	"log"

	"github.com/scalent-io/shiprocket/auth"
)

func main() {
	config, err := InitConfig()
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	token, err := auth.GetToken(config.Config.BaseURL, config.Config.Email, config.Config.Password)
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	//  err = ini()
	// if err != nil {
	// 	log.Fatal(err)
	// 	fmt.Println("err", err)
	// }

	fmt.Println("token----", token)
	fmt.Println("config :-----", config)
}

// func GetShiprocketToken(ctx context.Context, shiprocketConfig ShiprocketConfig) (string, error) {

// 	config, err := InitConfig()
// 	if err != nil {
// 		log.Fatal(err)
// 		fmt.Println("err", err)
// 	}

// 	token, err := auth.GetToken(config.Config.BaseURL, config.Config.Email, config.Config.Password)
// 	if err != nil {
// 		log.Fatal(err)
// 		fmt.Println("err", err)
// 	}

// 	return token, nil
// }
