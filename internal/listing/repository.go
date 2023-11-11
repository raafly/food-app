/*
This repository is divided into three main codes namely customer, product, cart.
*/

package listing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/raafly/food-app/pkg/helpers"
)

// Customer

type CustomerRepository interface {
	CreateAccount(ctx context.Context, tx *sql.Tx, model Customers) Customers
	LoginAccount(ctx context.Context, db *sql.DB, model Customers) (Customers, error)
	UpdatePhone(ctx context.Context, tx *sql.Tx, model Customers)
	UpdateAddress(ctx context.Context, tx *sql.Tx, model Customers)
	FindById(ctx context.Context, db *sql.DB, userId string) (Customers, error)
}

type CustomerRepositoryImpl struct {
}

func NewCustomerRepository() CustomerRepository {
	return &CustomerRepositoryImpl{}
}

func (repo *CustomerRepositoryImpl) CreateAccount(ctx context.Context, tx *sql.Tx, model Customers) Customers {
	SQL := "INSERT INTO customers(id, username, email, password) VALUES($1, $2, $3, $4)"
	_, err := tx.ExecContext(ctx, SQL, model.Id, model.Username, model.Email, model.Password)
	if err != nil {
		log.Println("error in line 33", err)
	}
	/*
	INSERT INTO cart (user_id) VALUES (123);
	DONT FORGET TO CREATE A NEW CART IF USER DOENS NOT EXIT CART
	*/
	return model
}

func (repo *CustomerRepositoryImpl) LoginAccount(ctx context.Context, db *sql.DB, model Customers) (Customers, error) {
	SQL := "SELECT username, email, password FROM customers WHERE email = $1"
	rows, err := db.QueryContext(ctx, SQL, model.Email)
	helper.PanicIfError(err)

	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&model.Username, &model.Email, &model.Password)
		helper.PanicIfError(err)
		return model, nil
	}

	return model, errors.New("user not found")
}

func (repo *CustomerRepositoryImpl) UpdatePhone(ctx context.Context, tx *sql.Tx, model Customers) {
	SQL := "UPDATE customers SET phone = $1 WHERE username = $2"
	_, err := tx.ExecContext(ctx, SQL, model.Phone, model.Username)
	helper.PanicIfError(err)
}

func (repo *CustomerRepositoryImpl) UpdateAddress(ctx context.Context, tx *sql.Tx, model Customers) {
	SQL := "UPDATE customers SET address = $1 WHERE username = $2"
	_, err := tx.ExecContext(ctx, SQL, model.Address, model.Username)
	helper.PanicIfError(err)
}	

func (repo *CustomerRepositoryImpl) FindById(ctx context.Context, db *sql.DB, userId string) (Customers, error) {
	SQL := "SELECT id, username, email FROM customers WHERE username = $1"
	rows, err := db.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	customer := Customers{}
	if rows.Next() {
		err := rows.Scan(&customer.Id, &customer.Username, &customer.Email)
		helper.PanicIfError(err)
		return customer, nil
	} else {
		return customer, errors.New("account not found")
	}
}

// Product

type ProductRepository interface {
	ProductCreate(ctx context.Context, tx *sql.Tx, model Products) Products
	ProductUpdate(ctx context.Context, tx *sql.Tx, model ModelProductUpdate) ModelProductUpdate
	ProductDelete(ctx context.Context, tx *sql.Tx, productName string)
}

type ProductRepositoryImpl struct {
}

func NewProductRepository() ProductRepository {
	return &ProductRepositoryImpl{}
}

func (repo *ProductRepositoryImpl) ProductCreate(ctx context.Context, tx *sql.Tx, product Products) Products {
	log.Println(product)
	SQL := "INSERT INTO products(name, description, quantity, price, category) VALUES($1, $2, $3, $4, $5) RETURNING id"
	var id int
	_ = tx.QueryRowContext(ctx, SQL, product.Name, product.Description, product.Quantity, product.Price, product.Category).Scan(&id)	
	product.Id = int(id)

	return product
}

func (repo *ProductRepositoryImpl) ProductUpdate(ctx context.Context, tx *sql.Tx, product ModelProductUpdate) ModelProductUpdate {
	SQL := "UPDATE products SET quantity = $1 WHERE name = $2"
	_, err := tx.ExecContext(ctx, SQL, product.Quantity, product.Name)
	helper.PanicIfError(err)

	return product
}

func (repo *ProductRepositoryImpl) ProductDelete(ctx context.Context, tx *sql.Tx, productName string) {
	SQL := "DELETE FROM products WHERE name = $1"
	_, err := tx.ExecContext(ctx, SQL, productName)
	helper.PanicIfError(err)
}

// cart

type CartRepository interface {
	AddItem(data CartsDetail) error
	RemoveItem(id int) error
	GetAllCart(cartId int) ([]CartsDetail, error)
}

type CartRepositoryImpl struct {
	db *sql.DB
}

func NewCartRepository(db *sql.DB) CartRepository {
	return &CartRepositoryImpl{db: db}
}

func (r CartRepositoryImpl) AddItem(data CartsDetail) error {
	SQL := "INSERT INTO carts_detail(cart_id, product_id, quantity) VALUES($1, $2, $3)"
	if _, err := r.db.Exec(SQL, data.CartId, data.ProductId, data.Quantity); err != nil {
		return fmt.Errorf("failed to exec query %v", err)
	}

	return nil
}

func (r CartRepositoryImpl) RemoveItem(id int) error {
	SQL := `delete from carts_detail where id = $1`

	if _, err := r.db.Exec(SQL, id); err != nil {
		return fmt.Errorf("item id not found %v", err)
	}
	return nil
}

func (r CartRepositoryImpl) GetAllCart(cartId int) ([]CartsDetail, error) {
	SQL := `SELECT carts.id, carts_detail.id, carts_detail.product_id, carts_detail.price, carts_detail.quantity FROM carts_detail
	JOIN carts ON carts.id=carts_detail.cart_id WHERE cart_id = $1`

	if rows, err := r.db.Query(SQL, cartId); err != nil {
		return nil, fmt.Errorf("failed exec query %v", err)
	} else {
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
}