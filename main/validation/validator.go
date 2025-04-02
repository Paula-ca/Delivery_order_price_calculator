package validation

import (
	"example/DOPC/main/entities"
	"fmt"
)

// function to validate small_order_surcharge not be negative
func ValidateSmallOrderSurcharge(small_order_surcharge int) int {
	if small_order_surcharge < 0 {
		small_order_surcharge = 0
	}
	return small_order_surcharge
}

// function to validate distance
func ValidateDistanceHandler(distance int, distance_ranges *entities.DistanceRanges) error {

	if distance >= 0 && distance < 500 {
		distance_ranges.A = 0
		distance_ranges.B = 0.0
	} else if distance >= 500 && distance < 1000 {
		distance_ranges.A = 100
		distance_ranges.B = 0.0
	} else if distance >= 1000 && distance < 1500 {
		distance_ranges.A = 200
		distance_ranges.B = 0.0
	} else if distance >= 1500 && distance < 2000 {
		distance_ranges.A = 200
		distance_ranges.B = 1.0
	} else {
		return fmt.Errorf("the delivery is currently not available for the distance: %d meters", distance)
	}
	return nil
}

// function to validate cart price
func ValidateCartPrice(price int) error {
	if price < 0 {
		return fmt.Errorf("error, invalid cart value")
	}

	return nil
}
func ValidateVenueSlug(venue string) error {
	var arr [4]string

	arr[0] = "home-assignment-venue-helsinki"
	arr[1] = "home-assignment-venue-berlin"
	arr[2] = "home-assignment-venue-tokyo"
	arr[3] = "home-assignment-venue-stockholm"

	for _, item := range arr {
		if item == venue {
			return nil
		}
	}
	err := fmt.Errorf("error, invalid venue_slug provided")
	return err
}
