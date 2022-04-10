package cart

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/transactions"
	transaction2 "github.com/ramdanariadi/grocery-be-golang/main/services/transaction"
	"net/http"
	"strconv"
)

type CartControllerImpl struct {
	Service transaction2.CartService
}

func NewCartController(db *sql.DB) *CartControllerImpl {
	return &CartControllerImpl{
		Service: transaction2.CartServiceImpl{
			CartRepository: transactions.CartRepositoryImpl{
				DB: db,
			},
			ProductRepository: product.ProductRepositoryImpl{
				DB: db,
			},
		},
	}
}

func (controller CartControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	carts := controller.Service.FindById(id)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller CartControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	total, err := strconv.Atoi(vars["total"])
	helpers.PanicIfError(err)

	code := http.StatusInternalServerError
	if controller.Service.Save(userId, productId, total) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller CartControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	code := http.StatusNotModified
	if controller.Service.Delete(userId, productId) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}
