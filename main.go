package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/raafly/food-app/config"
	"github.com/raafly/food-app/src/listing"
)
 
func main() {
	db := config.NewDB()
	validate := validator.New()

	customerRepo := listing.NewCustomerRepository()
	customerSer := listing.NewCustomerService(customerRepo, db, validate)
	customerHandler := listing.NewCustomerHandler(customerSer)

	productRepo := listing.NewProductRepository()
	productSer := listing.NewProductService(productRepo, db, validate)
	productHandler := listing.NewProductHandler(productSer)
	route := listing.Route(customerHandler, productHandler)

	server := http.Server {
		Handler: route,
		Addr: ":3000",
	}

	server.ListenAndServe()
}