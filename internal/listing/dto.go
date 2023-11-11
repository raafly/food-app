package listing

// model customers

type ModelCustomerSignUp struct {
	Username	string	`json:"username" validate:"required,alpha,min=6"`
	Email		string	`json:"email" validate:"required,email"`
	Passsword	string	`json:"password" validate:"required,min=8"`
}

type ModelCustomerSignIn struct {
	Email		string	`json:"email"`
	Passsword	string	`json:"password"`
}

type ModelCustomerResponse struct {
	Id			string	`json:"id"`
	Username	string	`json:"username"`
	Email		string	`json:"email"`
}

type ModelCustomerUpdate struct {
	Username	string	`json:"username" validate:"required"`
	Phone		string	`josn:"phone"`
	Address		string	`json:"address"`
}

// model products

type ModelProductsCreate struct {
	Name		 string		`json:"name" validate:"required"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity" validate:"required"`
	Price		 int		`json:"price" validate:"required"`
	Category	 string		`json:"category" validate:"required"`
}

type ModelProductUpdate struct {
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Price		 int		`json:"price"`
	Quantity	 int		`json:"quantity"`
}

type ModelProductResponse struct {
	Id			 int		`json:"id"`
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 int		`json:"price"`
	Category	 string		`json:"category"`
}

// model cart

type CategoryReq struct {
	UserId	string	`json:"userId"`
}

type AddToCart struct {
	CartId		int	`json:"cartId" validate:"required"`
	ProductId	int	`json:"productId"`
	Quantity	int	`json:"quantity"`
}

type RemoveItemCart struct {
	CartDetailId	int	`json:"cartDetailId" validate:"required"`
	ProductId		int	`json:"productId" validate:"required"`
}