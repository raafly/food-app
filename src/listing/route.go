package listing

import (
	"github.com/julienschmidt/httprouter"
)

func Route(customer CustomerHandler, product ProductHandler) *httprouter.Router{
	route := httprouter.New()

	route.POST("/api/users/signup", customer.SignUp)
	route.POST("/api/users/signin", customer.SignIn)
	route.PUT("/api/users/phone", customer.UpdatePhone)
	route.PUT("/api/users/address", customer.UpdateAddress)
	route.GET("/api/users/{userId}", customer.FindById)

	route.POST("/api/products/", product.Create)
	route.PUT("/api/products/productName", product.Update)
	route.DELETE("/api/products/productName", product.Delete)

	return route
}