package listing

import (
	"context"
	"database/sql"
	"errors"

	"github.com/raafly/food-app/helper"
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
	helper.PanicIfError(err)

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
	SQL := "INSERT INTO products(id, name, description, quantity, price, category) VALUES(?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, product.Id, product.Name, product.Description, product.Quantity, product.Price, product.Category)
	if err != nil {
		panic(err)
	}

	_, err = result.LastInsertId()
	if err != nil {
		panic(err)
	}

	return product
}

func (repo *ProductRepositoryImpl) ProductUpdate(ctx context.Context, tx *sql.Tx, product ModelProductUpdate) ModelProductUpdate {
	SQL := "UPDATE products SET quantity = $1 WHERE name = $2"
	_, err := tx.ExecContext(ctx, SQL, product.Quantity, product.Name)

	if err != nil {
		panic(err)
	}
	return product
}

func (repo *ProductRepositoryImpl) ProductDelete(ctx context.Context, tx *sql.Tx, productName string) {
	SQL := "DELETE FROM products WHERE name = $1"
	_, err := tx.ExecContext(ctx, SQL, productName)
	if err != nil {
		panic(err)
	}
}
