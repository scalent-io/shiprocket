package shiprocket

type ClientResponse struct {
	Token string `json:"token"`
}

type CourierAvailabityResponse struct {
	CompanyAutoShipmentInsuranceSetting bool `json:"company_auto_shipment_insurance_setting"`
	CovidZones                          struct {
		DeliveryZone interface{} `json:"delivery_zone"`
		PickupZone   interface{} `json:"pickup_zone"`
	} `json:"covid_zones"`
	Currency string `json:"currency"`
	Data     struct {
		AvailableCourierCompanies []struct {
			AirMaxWeight           string        `json:"air_max_weight"`
			AssuredAmount          int           `json:"assured_amount"`
			BaseCourierID          interface{}   `json:"base_courier_id"`
			BaseWeight             string        `json:"base_weight"`
			Blocked                int           `json:"blocked"`
			CallBeforeDelivery     string        `json:"call_before_delivery"`
			ChargeWeight           float64       `json:"charge_weight"`
			City                   string        `json:"city"`
			Cod                    int           `json:"cod"`
			CodCharges             int           `json:"cod_charges"`
			CodMultiplier          int           `json:"cod_multiplier"`
			Cost                   string        `json:"cost"`
			CourierCompanyID       int           `json:"courier_company_id"`
			CourierName            string        `json:"courier_name"`
			CourierType            string        `json:"courier_type"`
			CoverageCharges        int           `json:"coverage_charges"`
			CutoffTime             string        `json:"cutoff_time"`
			DeliveryBoyContact     string        `json:"delivery_boy_contact"`
			DeliveryPerformance    float64       `json:"delivery_performance"`
			Description            string        `json:"description"`
			Edd                    string        `json:"edd"`
			EntryTax               int           `json:"entry_tax"`
			EstimatedDeliveryDays  string        `json:"estimated_delivery_days"`
			Etd                    string        `json:"etd"`
			EtdHours               int           `json:"etd_hours"`
			FreightCharge          float64       `json:"freight_charge"`
			ID                     int           `json:"id"`
			IsCustomRate           int           `json:"is_custom_rate"`
			IsHyperlocal           bool          `json:"is_hyperlocal"`
			IsInternational        int           `json:"is_international"`
			IsRtoAddressAvailable  bool          `json:"is_rto_address_available"`
			IsSurface              bool          `json:"is_surface"`
			LocalRegion            int           `json:"local_region"`
			Metro                  int           `json:"metro"`
			MinWeight              float64       `json:"min_weight"`
			Mode                   int           `json:"mode"`
			Odablock               bool          `json:"odablock"`
			OtherCharges           int           `json:"other_charges"`
			Others                 string        `json:"others"`
			PickupAvailability     string        `json:"pickup_availability"`
			PickupPerformance      float64       `json:"pickup_performance"`
			PickupPriority         string        `json:"pickup_priority"`
			PickupSupressHours     int           `json:"pickup_supress_hours"`
			PodAvailable           string        `json:"pod_available"`
			Postcode               string        `json:"postcode"`
			QcCourier              int           `json:"qc_courier"`
			Rank                   string        `json:"rank"`
			Rate                   float64       `json:"rate"`
			Rating                 float64       `json:"rating"`
			RealtimeTracking       string        `json:"realtime_tracking"`
			Region                 int           `json:"region"`
			RtoCharges             float64       `json:"rto_charges"`
			RtoPerformance         float64       `json:"rto_performance"`
			SecondsLeftForPickup   int           `json:"seconds_left_for_pickup"`
			SecureShipmentDisabled bool          `json:"secure_shipment_disabled"`
			ShipType               int           `json:"ship_type"`
			State                  string        `json:"state"`
			SuppressDate           string        `json:"suppress_date"`
			SuppressText           string        `json:"suppress_text"`
			SuppressionDates       []interface{} `json:"suppression_dates"`
			SurfaceMaxWeight       string        `json:"surface_max_weight"`
			TrackingPerformance    float64       `json:"tracking_performance"`
			VolumetricMaxWeight    interface{}   `json:"volumetric_max_weight"`
			WeightCases            float64       `json:"weight_cases"`
			Zone                   string        `json:"zone"`
		} `json:"available_courier_companies"`
		ChildCourierID            interface{} `json:"child_courier_id"`
		IsRecommendationEnabled   int         `json:"is_recommendation_enabled"`
		RecommendationAdvanceRule int         `json:"recommendation_advance_rule"`
		RecommendedBy             struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"recommended_by"`
		RecommendedCourierCompanyID    interface{} `json:"recommended_courier_company_id"`
		ShiprocketRecommendedCourierID interface{} `json:"shiprocket_recommended_courier_id"`
	} `json:"data"`
	DgCourier                    int           `json:"dg_courier"`
	EligibleForInsurance         int           `json:"eligible_for_insurance"`
	InsuraceOptedAtOrderCreation bool          `json:"insurace_opted_at_order_creation"`
	IsAllowTemplatizedPricing    bool          `json:"is_allow_templatized_pricing"`
	IsLatlong                    int           `json:"is_latlong"`
	IsOldZoneOpted               bool          `json:"is_old_zone_opted"`
	IsZoneFromMongo              bool          `json:"is_zone_from_mongo"`
	LabelGenerateType            int           `json:"label_generate_type"`
	OnNewZone                    int           `json:"on_new_zone"`
	SellerAddress                []interface{} `json:"seller_address"`
	Status                       int           `json:"status"`
	UserInsuranceManadatory      bool          `json:"user_insurance_manadatory"`
}
