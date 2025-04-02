## Delivery order price calculator service (DOPC)
DOPC is an imaginary backend service which is capable of calculating the total price and price breakdown of a delivery order. 
DOPC integrates with the Home Assignment API to fetch venue related data required to calculate the prices. The term venue refers to any kind of restaurant / shop / store.
The DOPC service should provide a single endpoint: GET /api/v1/delivery-order-price, which takes the following as query parameters (all are required):

    venue_slug (string): The unique identifier (slug) for the venue from which the delivery order will be placed
    cart_value: (integer): The total value of the items in the shopping cart
    user_lat (number with decimal point): The latitude of the user's location
    user_lon (number with decimal point): The longitude of the user's location

So, an example request to DOPC could look like this:
http://localhost:8080/api/v1/delivery-order-price?venue_slug=home-assignment-venue-helsinki&cart_value=1000&user_lat=60.17094&user_lon=24.93087

The endpoint should return a JSON response in the following format:

{
  "total_price": 1190,
  "small_order_surcharge": 0,
  "cart_value": 1000,
  "delivery": {
    "fee": 190,
    "distance": 177
  }
}

Where:

    total_price (integer): The calculated total price
    small_order_surcharge (integer): The calculated small order surcharge
    cart_value (integer): The cart value. This is the same as what was got as query parameter.
    delivery (object): An object containing:
        fee (integer): The calculated delivery fee
        distance (integer): The calculated delivery distance in meters
# Postman Documantation
https://documenter.getpostman.com/view/18629048/2sB2cSfiDS
