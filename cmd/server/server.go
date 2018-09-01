package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	service "github.com/paksuy/golang-tutor/webservice3"
)

func main() {
	router := httprouter.New()

	router.GET("/products/:product_id", service.GetProducts)

	http.ListenAndServe(":9000", router)
}
