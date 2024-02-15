package shiprocket

type CourierAvailabityRequest struct {
	PickupPostcode   int     `json:"pickup_postcode"`
	DeliveryPostcode int     `json:"delivery_postcode"`
	OrderID          *int    `json:"order_id"`
	Cod              int     `json:"cod"`
	Weight           float64 `json:"weight"`
}
