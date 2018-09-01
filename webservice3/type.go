package webservice3

// Product struct holds information about product id, product name and product price
// we also define its json form using tag
// to convert struct -> json and vice versa, golang use tags
// Product json will looks like:
// {
//		"product_id": 1,
//		"name": "laptop one",
//		"price": 12.5
// }
type Product struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

// User struct holds information about user id, user name and user address
// User json will looks like:
// {
//		"user_id": 1,
//		"name": "John Doe",
//		"address": "5th avenue, Grom City"
// }
type User struct {
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

// Order struct holds information about order id, user data and product data
// Order json will looks like:
// {
//		"order_id": 1,
//		"customer": {
//					"user_id": 1,
//					"name": "John Doe",
//					"address": "5th avenue, Grom City"
//			},
//		"items": [
//					{
//							"product_id": 1,
//							"name": "laptop one"
//							"price": 12.5
//					},
//					{
//							"product_id": 21,
//							"name": "power cord"
//							"price": 1.75
//					}
//			]
// }
type Order struct {
	OrderID  int64     `json:"order_id"`
	Customer User      `json:"customer"`
	Items    []Product `json:"items"`
}
