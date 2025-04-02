package entities

// RESPONSE STRUCTURE
type OrderPrice struct {
	TotalPrice          int      `json:"total_price"`
	SmallOrderSurcharge int      `json:"small_order_surcharge"`
	CartValue           int      `json:"cart_value"`
	Delivery            Delivery `json:"delivery"` // Nested struct Delivery
}
type Delivery struct {
	Fee      int `json:"fee"`
	Distance int `json:"distance"`
}

// STATIC API
type StaticResponse struct {
	VenueRaw VenueRawS `json:"venue_raw"`
}
type VenueRawS struct {
	Location Location `json:"location"`
}

type Location struct {
	Coordinates []float64 `json:"coordinates"`
}

// DYNAMIC API
type DynamicResponse struct {
	VenueRaw VenueRawD `json:"venue_raw"`
}
type VenueRawD struct {
	Delivery_specs DeliverySpecs `json:"delivery_specs"`
}
type DeliverySpecs struct {
	Order_minimum_no_surcharge int             `json:"order_minimum_no_surcharge"`
	DeliveryPricing            DeliveryPricing `json:"delivery_pricing"`
}

type DeliveryPricing struct {
	BasePrice      int              `json:"base_price"`
	DistanceRanges []DistanceRanges `json:"distance_ranges"`
}
type DistanceRanges struct {
	Min int     `json:"min"`
	Max int     `json:"max"`
	A   int     `json:"a"`
	B   float32 `json:"b"`
}

func NewOrderPrice(total_price int, small_order_surcharge int, cart_value int, fee int, distance int) *OrderPrice {
	return &OrderPrice{
		TotalPrice:          total_price,
		SmallOrderSurcharge: small_order_surcharge,
		CartValue:           cart_value,
		Delivery: Delivery{
			Fee:      fee,
			Distance: distance,
		},
	}
}
