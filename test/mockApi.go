package test

import "example/DOPC/main/api"

// Mock API Client for testing
type MockAPIClient struct{}

func (m *MockAPIClient) CallStaticAPI(url string, lat float64, lon float64) api.ApiError {
	return nil // Simulate successful call
}

func (m *MockAPIClient) CallDynamicAPI(url string, cartValue int) (interface{}, api.ApiError) {
	return map[string]string{"status": "success"}, nil // Simulate successful response
}

// Mock API Client for testing, simulating an error
type MockAPIClientWithError struct{}

func (m *MockAPIClientWithError) CallStaticAPI(url string, lat float64, lon float64) api.ApiError {
	return api.NewApiError("Static API error", 500) // Simulate an error
}

func (m *MockAPIClientWithError) CallDynamicAPI(url string, cartValue int) (interface{}, api.ApiError) {
	return nil, nil // Simulate success
}

// Mock API Client for testing, simulating an error for Dynamic API
type MockAPIClientWithErrorDynamic struct{}

func (m *MockAPIClientWithErrorDynamic) CallStaticAPI(url string, lat float64, lon float64) api.ApiError {
	return nil // Simulate success for Static API
}

func (m *MockAPIClientWithErrorDynamic) CallDynamicAPI(url string, cartValue int) (interface{}, api.ApiError) {
	return nil, api.NewApiError("Dynamic API error", 500) // Simulate an error for Dynamic API
}
