package webservice3

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

var userList []User
var productList []Product
var orderList []Order

func init() {
	// in go, we can create/instantiate a struct like following statement. It's called composite statement.
	user := User{UserID: 1, Name: "John Doe", Address: "5th avenue, Grom City"}

	// in go, the following statement appends a data into a slice (list/array)
	userList = append(userList, user)

	// alternatively, we can instantiate a struct using empty composite.
	product := Product{}

	// and then fill each field afterwards
	product.ProductID = 1
	product.Name = "Laptop One"
	product.Price = 12.5

	// appending product to productList
	productList = append(productList, product)
}

type productListResponse struct {
	Data []Product `json:"data"`
}

// GetProducts handles GET request. It takes an in-url parameter. It can be "list" or
// numeric product id
func GetProducts(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	param := p.ByName("product_id")

	if param == "list" {
		resp := productListResponse{Data: productList}
		respByte, err := json.Marshal(&resp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(respByte)
		return
	}

	prdID, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, p := range productList {
		if p.ProductID == prdID {
			respByte, err := json.Marshal(&p)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(respByte)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
