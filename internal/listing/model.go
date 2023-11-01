package listing

// model database 

type Customers struct {
	Id				string
	Username		string
	Email			string
	Password		string
	Phone			string
	Address			string
}

type Products struct {
	Id			 int		
	Name		 string	  
	Description	 string		
	Quantity	 int 
	Price		 uint
	Category	 string
}

// customer model request and response

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

// product model request and response

type ModelProductsCreate struct {
	Name		 string		`json:"name" validate:"required"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity" validate:"required"`
	Price		 uint		`json:"price" validate:"required"`
	Category	 string		`json:"category" validate:"required"`
}

type ModelProductUpdate struct {
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Price		 uint		`json:"price"`
	Quantity	 int		`json:"quantity"`
}

type ModelProductResponse struct {
	Id			 int		`json:"id"`
	Name		 string		`json:"name"`
	Description	 string		`json:"description"`
	Quantity	 int		`json:"quantity"`
	Price		 uint		`json:"price"`
	Category	 string		`json:"category"`
}

// web response 

type WebResponse struct {
	Code		int
	Status		string
	Data		interface{}
	Err			interface{}
}