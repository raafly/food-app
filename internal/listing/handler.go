package listing

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/raafly/food-app/pkg/helpers"
)

// custoomer

type CustomerHandler interface {
	SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdatePhone(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	UpdateAddress(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type CustomerHandlerImpl struct {
	Port CustomerService	
}

func NewCustomerHandler(port CustomerService) CustomerHandler {
	return &CustomerHandlerImpl {
		Port: port,
	}
}

func (handler *CustomerHandlerImpl) SignUp(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelCustomerSignUp{}
	helper.ReadFromRequestBody(r, &request)

	data := handler.Port.SignUp(r.Context(), request)
	response := WebResponse {
		Code: 201,
		Status: "CREATED",
		Data: data,
		Err: nil,
	}

	helper.WriteToRequestBody(w, response)
}

func (handler *CustomerHandlerImpl) SignIn(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelCustomerSignIn{}
	helper.ReadFromRequestBody(r, &request)

	data , err := handler.Port.SingIn(r.Context(), request)
	helper.PanicIfError(err)

	response := WebResponse {
		Code: 200,
		Status: "OK",
		Data: data,
	}

	helper.WriteToRequestBody(w, response)
}

func (handler *CustomerHandlerImpl)	UpdatePhone(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelCustomerUpdate{}
	helper.ReadFromRequestBody(r, &request)

	handler.Port.UpdatePhone(r.Context(), request)
	response := WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)
}

func (handler *CustomerHandlerImpl)	UpdateAddress(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelCustomerUpdate{}
	helper.ReadFromRequestBody(r, &request)

	handler.Port.UpdateAddress(r.Context(), request)
	response := WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)
}

func (handler *CustomerHandlerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	userId := params.ByName("userId")

	data, err := handler.Port.FindById(r.Context(), userId)
	helper.PanicIfError(err)

	response := WebResponse {
		Code: 200,
		Status: "OK",
		Data: data,
	}

	helper.WriteToRequestBody(w, response)
}

// product

type ProductHandler interface {
	Create(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}

type ProductHandlerImpl struct {
	Port 	ProductService
}

func NewProductHandler(port ProductService) ProductHandler {
	return &ProductHandlerImpl {
		Port: port,
	}
}

func(handler *ProductHandlerImpl) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelProductsCreate{}
	helper.ReadFromRequestBody(r, &request) 

	data := handler.Port.Create(r.Context(), request)
	response := WebResponse {
		Code: 201,
		Status: "CREATED",
		Data: data,
	}

	helper.WriteToRequestBody(w, response)
}
 
func(handler *ProductHandlerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	request := ModelProductUpdate{}
	helper.ReadFromRequestBody(r, &request) 

	productName := params.ByName("name")
	request.Name = productName

	data := handler.Port.Update(r.Context(), request)
	response := WebResponse {
		Code: 201,
		Status: "CREATED",
		Data: data,
	}

	helper.WriteToRequestBody(w, response)
}
 
func(handler *ProductHandlerImpl) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	productName := params.ByName("name")
	handler.Port.Delete(r.Context(), productName)
	response := WebResponse {
		Code: 200,
		Status: "OK",
	}

	helper.WriteToRequestBody(w, response)
}
