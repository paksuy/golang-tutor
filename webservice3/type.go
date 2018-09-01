package webservice3

type Product struct {
	ProductID int64   `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
}

type User struct {
	UserID  int64  `json:"user_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Order struct {
	OrderID  int64     `json:"order_id"`
	Customer User      `json:"customer"`
	Items    []Product `json:"items"`
}
