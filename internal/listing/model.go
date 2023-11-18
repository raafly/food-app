package listing

import "time"

type Customers struct {	
	Id, Username, Email, Password, Phone, Address string	
}

type Products struct {
	Id, Quantity, Price int		
	Name, Description, Category string			
}

type Carts struct {
	Id			int
	User_id		string
	Item		CartsDetail
	Created_at	time.Time
}

type CartsDetail struct {
	Id, CartId, ProductId, Quantity, Price int
}

type WebResponse struct {
	Code		int
	Status		string
	Data		interface{}
}