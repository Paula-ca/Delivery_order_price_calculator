# Delivery Order Price Calculator Service (DOPC)
## Overview

The Delivery Order Price Calculator (DOPC) is a Go backend service designed to calculate the total price and the breakdown of a delivery order. It integrates with the Home Assignment API to fetch venue-specific data necessary for calculating the prices.

The term venue refers to any kind of restaurant, shop, or store.

The DOPC service provides a single endpoint that accepts the following query parameters and returns a price breakdown for a delivery order.
## API Endpoint

Endpoint: ```GET /api/v1/delivery-order-price```

#### Query Parameters

The following query parameters are required for the endpoint:

|Parameter	|Type	    |Description                                                                                      |
|:---------:|:---------:|:-----------------------------------------------------------------------------------------------:|
|venue_slug	|string	    |A unique identifier (slug) for the venue (restaurant, shop, etc.)where the order will be placed. |
|cart_value	|integer	|The total value of the items in the shopping cart (in the currency unit).                        |
|user_lat	|float64	|The latitude of the user's location (decimal value).                                             |
|user_lon	|float64	|The longitude of the user's location (decimal value).                                            |

#### Example Request

An example of a request to the DOPC service could look like this:

```GET "http://localhost:8080/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087"```

In this example:

- The venue_slug is home-assignment-venue-helsinki.

- The cart_value is 1000 (total value of items in the shopping cart).

- The user_lat and user_lon are the latitude and longitude of the user's location in Helsinki.

## API Response

The response will be a JSON object that includes the calculated total price, the small order surcharge, the cart value, and the delivery details (fee and distance).

**Example Response:**
```json
{
  "total_price": 1190,
  "small_order_surcharge": 0,
  "cart_value": 1000,
  "delivery": {
    "fee": 190,
    "distance": 177
  }
}
```

## Response Fields
- total_price (integer): The calculated total price for the delivery order, which includes the cart value and the delivery fee.

- small_order_surcharge (integer): The surcharge applied if the order value is too small (e.g., a minimum cart value threshold). This value is 0 if no surcharge is applied.

- cart_value (integer): The cart value provided in the query parameters (same as cart_value).

- delivery (object):

  - fee (integer): The calculated delivery fee based on the venue and distance.

  - distance (integer): The calculated delivery distance in meters from the user's location to the venue.

## Error Handling

If any of the required query parameters are missing or invalid, or if there is an issue with the external API calls, the service will return an appropriate HTTP error response with a relevant error message.

For example:

  - If cart_value is not a valid integer, the response will be a 400 Bad Request with an error message like "Error, invalid cart_value".

  - If venue_slug is missing, the response will be a 400 Bad Request with an error message like "Error, invalid venue_slug provided".


## Postman Documentation

For a more detailed walkthrough and examples, you can refer to the Postman documentation for this API:

[View Postman Documentation](https://documenter.getpostman.com/view/18629048/2sB2cSfiDS).

**How It Works:**

  1. The user sends a GET request to the DOPC endpoint, providing the required query parameters (venue_slug, cart_value, user_lat, and user_lon).

  2. The service processes the request, validates the inputs, and then fetches data from the external Home Assignment API to calculate the delivery fee and distance.

  3. The service calculates the total price, including any applicable small order surcharges, and then sends the response in the specified JSON format.

**Notes:**

  - The DOPC service integrates with external APIs to fetch real-time venue data and calculate delivery charges. These external APIs may be subject to availability or changes in their structure.

  - The small_order_surcharge can be configured based on the venue’s business rules. By default, it is set to 0.

## Development & Testing

To run the DOPC service locally:

 **1. Clone the repository:**

```git clone https://github.com/Paula-ca/Delivery_order_price_calculator.git```

 **2. Navigate to the project directory:**

```cd DOPC```

 **3. Install the necessary dependencies:**

```go mod tidy```

 **4. Start the service:**

```go run main.go```

 **5. Open Postman or any other API client to send a GET request to the endpoint:**

  ```http://localhost:8080/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087```

  **6. You should receive a JSON response with the calculated price breakdown.**

## Conclusion

The Delivery Order Price Calculator (DOPC) service allows you to calculate the total price and price breakdown of a delivery order, including the delivery fee and distance. It integrates with external APIs to retrieve the necessary venue data and applies the relevant business rules to calculate the total price.


