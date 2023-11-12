/*
This service is divided into three main codes namely customer, product, cart.
*/

package listing

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/oklog/ulid/v2"
	config "github.com/raafly/food-app/pkg/configs"
	"github.com/raafly/food-app/pkg/exception"
	"github.com/raafly/food-app/pkg/helpers"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

// hash password

func hashedPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	return string(hash)
}

func matchPassword(passwordDb, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordDb), []byte(password))
	helper.PanicIfError(err)
	return true
}

// customer

type CustomerService interface {
	SignUp(ctx context.Context, request ModelCustomerSignUp) ModelCustomerResponse
	SingIn(ctx context.Context, request ModelCustomerSignIn) (ModelCustomerResponse, string, error)
	UpdatePhone(ctx context.Context, request ModelCustomerUpdate) 
	UpdateAddress(ctx context.Context, request ModelCustomerUpdate) 
	FindById(ctx context.Context, userId string) (ModelCustomerResponse, error)
}

type CustomerServiceImpl struct {
	Port		CustomerRepository
	DB 			*sql.DB
	Validate 	*validator.Validate
}

func NewCustomerService(port CustomerRepository, DB *sql.DB, validate *validator.Validate) CustomerService {
	return &CustomerServiceImpl{
		Port: port,
		DB: DB,
		Validate: validate,
	}
}

func (ser *CustomerServiceImpl) SignUp(ctx context.Context, request ModelCustomerSignUp) ModelCustomerResponse {
	err := ser.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	entropy := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	ms := ulid.Timestamp(time.Now())
	id, err := ulid.New(ms, entropy)
	helper.PanicIfError(err)

	hash := hashedPassword(request.Passsword)
	
	data := Customers {
		Id: id.String(),
		Username: request.Username,
		Email: request.Email,
		Password: hash,
	}
	
	response := ser.Port.CreateAccount(ctx, tx, data)
	helper.PanicIfError(err)

	return CustomerResponse(response)
}

func (ser *CustomerServiceImpl) SingIn(ctx context.Context, request ModelCustomerSignIn) (ModelCustomerResponse, string, error) {
	err  := ser.Validate.Struct(request)
	helper.PanicIfError(err)
	
	data := Customers {
		Email: request.Email,
		Password: request.Passsword,
	}
	
	data, err = ser.Port.LoginAccount(ctx, ser.DB, data)

	compare := matchPassword(data.Password, request.Passsword)
	if !compare {
		panic(exception.NewNotMatchError(err.Error()))
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Username: data.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "food-app",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	helper.PanicIfError(err)

	return CustomerResponse(data), token, nil
}

func (ser *CustomerServiceImpl)	UpdatePhone(ctx context.Context, request ModelCustomerUpdate) {
	err := ser.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	data := Customers {
		Username: request.Username,
		Phone: request.Phone,
	}
	

	ser.Port.UpdatePhone(ctx, tx, data)
}

func (ser *CustomerServiceImpl)	UpdateAddress(ctx context.Context, request ModelCustomerUpdate) {
	err := ser.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	data := Customers {
		Username: request.Username,
		Phone: request.Address,
	}
	
	ser.Port.UpdateAddress(ctx, tx, data)
}

func (ser *CustomerServiceImpl)	FindById(ctx context.Context, userId string) (ModelCustomerResponse, error) {
	customer, err := ser.Port.FindById(ctx, ser.DB, userId)
	helper.PanicIfError(err)

	return CustomerResponse(customer), nil
}

// Product

type ProductService interface {
	Create(ctx context.Context, request ModelProductsCreate) ModelProductResponse
	Update(ctx context.Context, request ModelProductUpdate) ModelProductUpdate
	Delete(ctx context.Context, productName string)
}

type ProductServiceImpl struct {
	ProductRepository ProductRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewProductService(productRepository ProductRepository, DB *sql.DB, validate *validator.Validate) ProductService {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		DB: DB,
		Validate: validate,
	}
}

func (ser *ProductServiceImpl) Create(ctx context.Context, request ModelProductsCreate) ModelProductResponse {
	err := ser.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	product := Products {
		Name: request.Name,
		Description: request.Description,
		Quantity: request.Quantity,
		Price: request.Price,
		Category: request.Category,
	}
	log.Println(product)	
	
	response := ser.ProductRepository.ProductCreate(ctx, tx, product)
	return ToProductReponse(response)
}

func (ser *ProductServiceImpl) Update(ctx context.Context, request ModelProductUpdate) ModelProductUpdate {
	err := ser.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	request = ModelProductUpdate {
		Quantity: request.Quantity,
	}
	product := ser.ProductRepository.ProductUpdate(ctx, tx, request)
	return ToProductReponseUpdate(product)
}

func (ser *ProductServiceImpl) Delete(ctx context.Context, productName string) {
	tx, err := ser.DB.Begin()
	defer helper.CommitOrRollback(tx)
	helper.PanicIfError(err)

	ser.ProductRepository.ProductDelete(ctx, tx, productName)
}

// carts

type CartService interface {
	CartAddItem(request AddToCart) error
	CartRemoveItem(request RemoveItemCart) error
	GetAllItem(cartId int) []ResCartsDetail
}

type CartServiceImpl struct {
	Port 		CartRepository
	DB 			*sql.DB
	Validate	*validator.Validate
}

func NewCartService(port CartRepository, DB *sql.DB, validate *validator.Validate) CartService {
	return &CartServiceImpl{
		Port: port,
		DB: DB,
		Validate: validate,
	}
}

func (s CartServiceImpl) CartAddItem(request AddToCart) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	data := CartsDetail {
		CartId: request.CartId,
		ProductId: request.ProductId,
		Quantity: request.Quantity,
	}

	if err := s.Port.AddItem(data); err != nil {
		return err 
	}
	return nil
}

func (s CartServiceImpl) GetAllItem(cartId int) []ResCartsDetail {
	cartDetail, err := s.Port.GetAllCart(cartId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return ToCartResponses(cartDetail)
}


func (s CartServiceImpl) CartRemoveItem(request RemoveItemCart) error {
	if err := s.Validate.Struct(request); err != nil {
		return err
	}

	if err := s.Port.RemoveItem(request.CartDetailId); err != nil {
		return err
	}
	return nil
}

