package api

import (
	"encoding/json"
	"example/DOPC/main/entities"
	"fmt"
	"io"
	"math"
	"net/http"
)

var Delivery_distance float64

// http GET to static API
func (r *RealAPIClient) CallStaticAPI(url string, usrlat, usrlong float64) ApiError {

	// GET request
	resp, err := http.Get(url)
	if err != nil {
		return NewApiError(fmt.Sprintf("error calling static API, this is the HTTP Status response from the API: %s. Go error: %s", resp.StatusCode, err.Error()), http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	// Read the body of the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return NewApiError("error decoding the API body into go struct", http.StatusInternalServerError)
	}
	// Parse to Go structure
	var result entities.StaticResponse
	if err := json.Unmarshal(body, &result); err != nil {
		apierr := NewApiError(fmt.Sprintf("error to parse to GO structure: %s. Go error: %s", err.Error()), http.StatusInternalServerError)
		return apierr
	}

	// Getters
	venueLat := result.VenueRaw.Location.Coordinates[1]
	venueLon := result.VenueRaw.Location.Coordinates[0]

	//Setters
	Delivery_distance = haversine(usrlat, usrlong, venueLat, venueLon)

	return nil

}
func DistanceExport() int {
	return int(Delivery_distance)
}

// haversine function to calculate the straight-line distance between two points
func haversine(lat1, lon1, lat2, lon2 float64) float64 {

	const earthRadius = 6371000 // Earth's radius in meters

	// Convert degrees to radians
	toRadians := func(degree float64) float64 {
		return degree * math.Pi / 180
	}
	lat1Rad, lon1Rad := toRadians(lat1), toRadians(lon1)
	lat2Rad, lon2Rad := toRadians(lat2), toRadians(lon2)

	// Haversine formula
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// Calculate distance
	return earthRadius * c
}
