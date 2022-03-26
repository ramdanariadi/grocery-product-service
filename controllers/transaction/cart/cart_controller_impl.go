package cart

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	"go-tunas/repositories/product"
	"go-tunas/repositories/transactions"
	"go-tunas/services/transaction"
	"net/http"
	"strconv"
)

type CartControllerImpl struct {
	Service transaction.CartService
}

func NewCartController(db *sql.DB) *CartControllerImpl {
	return &CartControllerImpl{
		Service: transaction.CartServiceImpl{
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
