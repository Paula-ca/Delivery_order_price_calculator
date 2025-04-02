package handler

import (
	"encoding/json"
	"example/DOPC/main/api"
	"example/DOPC/main/validation"
	"net/http"
	"strconv"
)

// Handler of the endpoint
func HandleRequest(w http.ResponseWriter, r *http.Request, apiClient api.APIClient) {
	//get the params of the query url
	queryParams := r.URL.Query()

	num, err := strconv.Atoi(queryParams.Get("cart_value"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	str := queryParams.Get("venue_slug")
	//function to validate venue slug string
	err = validation.ValidateVenueSlug(str)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lat, err := strconv.ParseFloat(queryParams.Get("user_lat"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lon, err := strconv.ParseFloat(queryParams.Get("user_lon"), 32)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//function to validate the cart price is not negative
	err = validation.ValidateCartPrice(num)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Call static API
	apierr := apiClient.CallStaticAPI("https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/"+str+"/static", lat, lon)
	if apierr != nil {
		errormsg, httpstatuserr := apierr.Error()
		http.Error(w, errormsg, httpstatuserr)
		return
	}
	//Call dynamic API
	response2, apierr := apiClient.CallDynamicAPI("https://consumer-api.development.dev.woltapi.com/home-assignment-api/v1/venues/"+str+"/dynamic", num)
	if apierr != nil {
		errormsg, httpstatuserr := apierr.Error()
		http.Error(w, errormsg, httpstatuserr)
		return
	}

	// Parse response to JSON and sent to client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response2)
}

// Create a wrapper function that calls HandleRequest with an injected APIClient
func HandleRequestWithClient(apiClient api.APIClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleRequest(w, r, apiClient)
	}
}
