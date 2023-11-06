package route

import (
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/food-app/internal/listing"
	"github.com/raafly/food-app/pkg/exception"
)

func NewRoute(customer listing.CustomerHandler, product listing.ProductHandler) *httprouter.Router{
	route := httprouter.New()

	route.POST("/api/users/signup", customer.SignUp)
	route.POST("/api/users/signin", customer.SignIn)
	route.PUT("/api/users/phone", customer.UpdatePhone)
	route.PUT("/api/users/address", customer.UpdateAddress)
	route.GET("/api/users/{userId}", customer.FindById)

	route.POST("/api/products/", product.Create)
	route.PUT("/api/products/name", product.Update)
	route.DELETE("/api/products/name", product.Delete)
	
	route.PanicHandler = exception.ErrorHadler

	return route
}