package test

import (
	"encoding/json"

	"example/DOPC/main/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test HandleRequest with Mock API Client
func TestHandleRequest_Valid(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cart_value=100&venue_slug=home-assignment-venue-helsinki&user_lat=40.7128&user_lon=-74.0060", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Use the Mock API client
	mockClient := &MockAPIClient{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequest(w, r, mockClient) // Inject the mock client into the handler
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the Content-Type header
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))

	// Check the body of the response
	var response map[string]string
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(response["status"])
	// Validate the response structure
	assert.Equal(t, "success", response["status"])
}

func TestHandleRequest_InvalidCartValue(t *testing.T) {
	// Test case where the cart_value is not a valid integer
	req, err := http.NewRequest("GET", "/?cart_value=invalid&venue_slug=valid_slug&user_lat=40.7128&user_lon=-74.0060", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Use the Mock API client
	mockClient := &MockAPIClient{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequest(w, r, mockClient) // Inject the mock client into the handler
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 400 Bad Request (due to invalid cart_value)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// Check the body of the response
	body := rr.Body.String()
	assert.Contains(t, body, "invalid syntax") // Ensure the error message contains "invalid syntax"
}

func TestHandleRequest_MissingVenueSlug(t *testing.T) {
	// Test case where venue_slug is missing
	req, err := http.NewRequest("GET", "/?cart_value=100&user_lat=40.7128&user_lon=-74.0060", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Use the Mock API client
	mockClient := &MockAPIClient{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequest(w, r, mockClient) // Inject the mock client into the handler
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 400 Bad Request (due to missing venue_slug)
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// Check the body of the response
	body := rr.Body.String()
	assert.Contains(t, body, "invalid venue_slug provided") // Ensure the error message contains the correct description
}

func TestHandleRequest_APIErrorStatic(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cart_value=100&venue_slug=valid_slug&user_lat=40.7128&user_lon=-74.0060", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Use the mock client that simulates an error
	mockClient := &MockAPIClientWithError{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequest(w, r, mockClient)
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 500 (Internal Server Error)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the body of the response
	body := rr.Body.String()
	assert.Contains(t, body, "Static API error") // Ensure the error message is returned
}
func TestHandleRequest_APIErrorDynamic(t *testing.T) {
	req, err := http.NewRequest("GET", "/?cart_value=100&venue_slug=valid_slug&user_lat=40.7128&user_lon=-74.0060", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Use the mock client that simulates an error for Dynamic API
	mockClient := &MockAPIClientWithErrorDynamic{}

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.HandleRequest(w, r, mockClient)
	})
	handler.ServeHTTP(rr, req)

	// Check if the status code is 500 (Internal Server Error)
	assert.Equal(t, http.StatusInternalServerError, rr.Code)

	// Check the body of the response
	body := rr.Body.String()
	assert.Contains(t, body, "Dynamic API error") // Ensure the error message is returned
}
