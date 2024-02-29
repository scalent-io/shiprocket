package shiprocket

func (s *ShiprocketClient) GetAvailableCouriers(request CourierAvailabityRequest) (*CourierAvailabityResponse, *ErrorResponse) {
	var response CourierAvailabityResponse
	err := SendRequest("GET", "/v1/external/courier/serviceability/", s.config.baseURL, s.tokenConfig.Token, request, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
