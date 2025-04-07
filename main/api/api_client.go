package api

// Define an interface for API calls
type APIClient interface {
	CallStaticAPI(url string, lat float64, lon float64) ApiError
	CallDynamicAPI(url string, cartValue int) (interface{}, ApiError)
}

// This is your real API client implementation.
type RealAPIClient struct{}
