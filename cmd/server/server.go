package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/paksuy/golang-tutor/webservice3"
)

func main() {
	router := httprouter.New()

	// julienschmidt's httprouter has a feature to receive data/value in url
	// the following router path will match /products/1, /products/2, /products/list etc
	router.GET("/products/:product_id", service.GetProducts)

	http.ListenAndServe(":9000", router)
}
