package auth

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/scalent-io/shiprocket/apimodel"
)

func GetToken(baseURL, email, password string) (string, error) {
	data := map[string]string{
		"email":    email,
		"password": password,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	resp, err := http.Post(baseURL+"/v1/external/auth/login", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var response apimodel.AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	return response.Token, err
}
