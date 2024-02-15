package shiprocket

func (s *ShiprocketClient) GetAvailableCouriers(request CourierAvailabityRequest) (*CourierAvailabityResponse, error) {
	// Create a new request
	resp, err := SendRequest("GET", "/v1/external/courier/serviceability/", s.config.BaseURL, s.tokenConfig.Token, request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var response CourierAvailabityResponse
	err = ReadResponse(resp, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
