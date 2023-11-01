package listing

// customer response 

func CustomerResponse(customer Customers) ModelCustomerResponse {
	return ModelCustomerResponse {
		Id: customer.Id,
		Username: customer.Username,
		Email: customer.Email,
	}
}

func CustomerResponsePhone(customer Customers) ModelCustomerUpdate {
	return ModelCustomerUpdate{
		Username: customer.Username,
		Phone: customer.Phone,
	}
}

func CustomerResponseAddress(customer Customers) ModelCustomerUpdate {
	return ModelCustomerUpdate{
		Username: customer.Username,
		Address: customer.Address,
	}
}

func ToProductReponse(product Products) ModelProductResponse {
	return ModelProductResponse{
		Id: product.Id,
		Name: product.Name,
		Description: product.Description,
		Quantity: product.Quantity,
		Price: product.Price,
		Category: product.Category,
	}
}

func ToProductReponseUpdate(product ModelProductUpdate) ModelProductUpdate {
	return ModelProductUpdate{
		Name: product.Name,
		Description: product.Description,
		Quantity: product.Quantity,
		Price: product.Price,
	}
}