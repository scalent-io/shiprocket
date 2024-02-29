package shiprocket

type CourierAvailabityRequest struct {
	PickupPostcode   int     `json:"pickup_postcode"`
	DeliveryPostcode int     `json:"delivery_postcode"`
	OrderID          *int    `json:"order_id"`
	Cod              int     `json:"cod"`
	Weight           float64 `json:"weight"`
}

type CreateOrderRequest struct {
	BillingAddress       string         `json:"billing_address"`
	BillingAddress2      string         `json:"billing_address_2"`
	BillingCity          string         `json:"billing_city"`
	BillingCountry       string         `json:"billing_country"`
	BillingCustomerName  string         `json:"billing_customer_name"`
	BillingEmail         string         `json:"billing_email"`
	BillingLastName      string         `json:"billing_last_name"`
	BillingPhone         string         `json:"billing_phone"`
	BillingPincode       string         `json:"billing_pincode"`
	BillingState         string         `json:"billing_state"`
	Breadth              float64        `json:"breadth"`
	ChannelID            string         `json:"channel_id"`
	Comment              string         `json:"comment"`
	GiftwrapCharges      int64          `json:"giftwrap_charges"`
	Height               float64        `json:"height"`
	Length               float64        `json:"length"`
	OrderDate            string         `json:"order_date"`
	OrderID              string         `json:"order_id"`
	OrderItems           []OrderDetails `json:"order_items"`
	PaymentMethod        string         `json:"payment_method"`
	PickupLocation       string         `json:"pickup_location"`
	ShippingAddress      string         `json:"shipping_address"`
	ShippingAddress2     string         `json:"shipping_address_2"`
	ShippingCharges      int64          `json:"shipping_charges"`
	ShippingCity         string         `json:"shipping_city"`
	ShippingCountry      string         `json:"shipping_country"`
	ShippingCustomerName string         `json:"shipping_customer_name"`
	ShippingEmail        string         `json:"shipping_email"`
	ShippingIsBilling    bool           `json:"shipping_is_billing"`
	ShippingLastName     string         `json:"shipping_last_name"`
	ShippingPhone        string         `json:"shipping_phone"`
	ShippingPincode      string         `json:"shipping_pincode"`
	ShippingState        string         `json:"shipping_state"`
	SubTotal             int64          `json:"sub_total"`
	TotalDiscount        int64          `json:"total_discount"`
	TransactionCharges   int64          `json:"transaction_charges"`
	Weight               float64        `json:"weight"`
}

type OrderDetails struct {
	Discount     string `json:"discount"`
	Hsn          int64  `json:"hsn"`
	Name         string `json:"name"`
	SellingPrice string `json:"selling_price"`
	Sku          string `json:"sku"`
	Tax          string `json:"tax"`
	Units        int64  `json:"units"`
}

type GenerateInvoiceRequest struct {
	Ids []int `json:"ids"`
}

type GenerateManifestRequest struct {
	ShipmentID []int `json:"shipment_id"`
}

type GenerateLabelRequest struct {
	ShipmentID []int `json:"shipment_id"`
}

type GenerateAWBRequest struct {
	CourierID  string `json:"courier_id"`
	ShipmentID string `json:"shipment_id"`
}
