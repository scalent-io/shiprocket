package shiprocket

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ShiprocketClient struct {
	config      ClientConfig
	tokenConfig TokenConfig
}

var shiprocketClient ShiprocketClient

func NewShiprocketClient(config ClientConfig) (*ShiprocketClient, error) {

	token, err := GetToken(config)
	if err != nil {
		return nil, err
	}

	config.baseURL = BASE_URL

	shiprocketClient = ShiprocketClient{
		config: config,
		tokenConfig: TokenConfig{
			Token:             token,
			CreatedOn:         time.Now(),
			NextRefreshOnTime: time.Now().Add((7 * 60) + 30*time.Minute), //token will updated after 7:30 hrs
		},
	}

	return &shiprocketClient, nil
}

func GetToken(config ClientConfig) (string, error) {
	data := map[string]string{
		"email":    config.Email,
		"password": config.Password,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	resp, err := http.Post(config.baseURL+"/v1/external/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response ClientResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Token, err
}
