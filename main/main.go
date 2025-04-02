package main

import (
	"example/DOPC/main/api"
	"example/DOPC/main/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create an instance of your real API client
	apiClient := &api.RealAPIClient{}
	// Define endpoint
	http.HandleFunc("/api/v1/delivery-order-price", handler.HandleRequestWithClient(apiClient))
	// Start up server
	fmt.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
