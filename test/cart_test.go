package test

import (
	"fmt"
	"testing"

	config "github.com/raafly/food-app/pkg/configs"
)


type CartsDetail struct {
	Id, CartId, ProductId, Quantity, Price int
}


func AddItem(data CartsDetail) error {
	db := config.NewDB()
	SQL := "INSERT INTO carts_detail(cart_id, product_id, quantity) VALUES($1, $2, $3)"
	if _, err := db.Exec(SQL, data.CartId, data.ProductId, data.Quantity); err != nil {
		return fmt.Errorf("failed to exec query %v", err)
	}
	return nil
}

func RemoveItem(id int) error {
	db := config.NewDB()
	SQL := `delete from carts_detail where id = $1`

	if _, err := db.Exec(SQL, id); err != nil {
		return fmt.Errorf("item id not found %v", err)
	}
	return nil
}

func FindCartById(id int) {
	
}

func GetCart() ([]CartsDetail, error) {
	db := config.NewDB()
	SQL := `SELECT carts.id, carts_detail.id, carts_detail.product_id, carts_detail.price, carts_detail.quantity FROM carts_detail
	JOIN carts ON carts.id=carts_detail.cart_id WHERE cart_id = 9`

	rows, err := db.Query(SQL)
	if err != nil {
		return nil, fmt.Errorf("failed exec query %v", err)
	}
	defer rows.Close()

	var carts []CartsDetail
	for rows.Next() {
		cart := CartsDetail{}
		err := rows.Scan(&cart.CartId, &cart.Id, &cart.ProductId, &cart.Price, &cart.Quantity)
		if err != nil {
			return nil, fmt.Errorf("failed to exec query %v", err)
		}
		carts = append(carts, cart)
	}

	return carts, nil
}

func TestAddItem(t *testing.T) {
	data := CartsDetail {
		CartId: 9,
		ProductId: 1,
		Quantity: 10,
	}

	if err := AddItem(data); err != nil {
		fmt.Println(err)	
	}
}

func TestGetAllCartItem(t *testing.T) {
	if carts, err := GetCart(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(carts)
	}
}

func TestRemoveItem(t *testing.T) {
	if err := RemoveItem(54); err != nil {
		fmt.Println(err)
	}
}