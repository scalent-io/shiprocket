package shiprocket

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func SendRequest(method, path string, BaseURL string, Token string, requestBody interface{}, responseBody interface{}) *ErrorResponse {

	var errResp ErrorResponse

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return &errResp
	}

	req, err := http.NewRequest(method, BaseURL+path, bytes.NewBuffer(jsonData))
	if err != nil {
		return &errResp
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &errResp
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return &errResp
	}

	if resp.StatusCode != http.StatusOK {
		var errResp ErrorResponse
		err = json.Unmarshal(respBody, &errResp)
		if err != nil {
			return &errResp
		}
		return &errResp
	}

	err = json.Unmarshal(respBody, &responseBody)
	if err != nil {
		return &errResp
	}

	return nil
}

// func SendRequest(method, path string, BaseURL string, Token string, body interface{}) (*http.Response, error) {
// 	jsonData, err := json.Marshal(body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest(method, BaseURL+path, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+Token)

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return resp, err
// }

// func ReadResponse(resp *http.Response, result interface{}) error {
// 	if resp.StatusCode != http.StatusOK {
// 		return fmt.Errorf("bad status code: %d", resp.StatusCode)
// 	}

// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	err = json.Unmarshal(body, result)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
