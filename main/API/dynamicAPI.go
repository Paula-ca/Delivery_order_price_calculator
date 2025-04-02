package api

import (
	"encoding/json"
	"example/DOPC/main/entities"
	"example/DOPC/main/validation"
	"fmt"
	"io"
	"net/http"
)

// http GET to dynamic API
func (r *RealAPIClient) CallDynamicAPI(url string, cartValue int) (interface{}, ApiError) {

	// GET request
	resp, err := http.Get(url)
	if err != nil {
		return entities.OrderPrice{}, NewApiError(fmt.Sprintf("error calling static API, this is the HTTP Status response from the API: %s. Go error: %s", resp.StatusCode, err.Error()), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.OrderPrice{}, NewApiError("error decoding the API body into go struct.", http.StatusInternalServerError)
	}
	// Parse to GO structure
	var result entities.DynamicResponse
	if err := json.Unmarshal(body, &result); err != nil {
		apierr := NewApiError(fmt.Sprintf("error to parse to go structure: %s. Go error: %s", err.Error()), http.StatusInternalServerError)
		return entities.OrderPrice{}, apierr
	}
	//getters
	order_minimum_no_surcharge := result.VenueRaw.Delivery_specs.Order_minimum_no_surcharge
	distance_ranges := result.VenueRaw.Delivery_specs.DeliveryPricing.DistanceRanges
	base_price := result.VenueRaw.Delivery_specs.DeliveryPricing.BasePrice

	//setters
	var smallOrderSurcharge = order_minimum_no_surcharge - cartValue
	var deliveryFee = base_price + distance_ranges[0].A + int(distance_ranges[0].B)*DistanceExport()/10
	var totalPrice = cartValue + smallOrderSurcharge + deliveryFee

	//validation of small_order_surcharge
	err = validation.ValidateSmallOrderSurcharge(smallOrderSurcharge)
	if err != nil {
		return entities.OrderPrice{}, NewApiError(err.Error(), http.StatusBadRequest)
	}
	//validation of distance
	err = validation.ValidateDistanceHandler(DistanceExport(), &distance_ranges[0])
	if err != nil {
		return entities.OrderPrice{}, NewApiError(err.Error(), http.StatusBadRequest)
	}
	newOrderPrice := entities.NewOrderPrice(totalPrice, smallOrderSurcharge, cartValue, deliveryFee, DistanceExport())

	return *newOrderPrice, nil

}
