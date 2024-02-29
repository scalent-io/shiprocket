package shiprocket

func (s *ShiprocketClient) CreateOrder(request CreateOrderRequest) (*CreateOrderResponse, *ErrorResponse) {
	var response CreateOrderResponse
	errResp := SendRequest("POST", "/v1/external/orders/create/adhoc", s.config.baseURL, s.tokenConfig.Token, request, &response)
	if errResp != nil {
		return nil, errResp
	}

	return &response, nil
}
