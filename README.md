## Delivery order price calculator service (DOPC)
DOPC is an imaginary backend service which is capable of calculating the total price and price breakdown of a delivery order. 
DOPC integrates with the Home Assignment API to fetch venue related data required to calculate the prices. The term venue refers to any kind of restaurant / shop / store.
The DOPC service should provide a single endpoint: GET /api/v1/delivery-order-price, which takes the following as query parameters (all are required):

    venue_slug (string): The unique identifier (slug) for the venue from which the delivery order will be placed
    cart_value: (integer): The total value of the items in the shopping cart
    user_lat (number with decimal point): The latitude of the user's location
    user_lon (number with decimal point): The longitude of the user's location
