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

func ToProductReponseUpdate(product ModelProductUpdate) ModelProductUpdate {
	return ModelProductUpdate{
		Name: product.Name,
		Description: product.Description,
		Quantity: product.Quantity,
		Price: product.Price,
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

func ToProductCart(cart CartsDetail) ResCartsDetail {
	return ResCartsDetail{
		Id: cart.Id,
		CartId: cart.CartId,
		ProductId: cart.ProductId,
		Quantity: cart.Quantity,
		Price: cart.Price,
	}
}

func ToCartResponses(cart []CartsDetail) []ResCartsDetail {
	var productResponses []ResCartsDetail
	for _, product := range cart {
		productResponses = append(productResponses, ToProductCart(product))
	}
	return productResponses
}