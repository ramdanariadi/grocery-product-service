package cart

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"go-tunas/customresponses"
	"go-tunas/repositories/product"
	"go-tunas/repositories/transactions"
	"go-tunas/services/transaction"
	"net/http"
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

func (controller CartControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("userId")
	carts := controller.Service.FindById(id)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller CartControllerImpl) Save(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	productId := param.ByName("productId")
	code := http.StatusInternalServerError
	if controller.Service.Save(userId, productId) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller CartControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	productId := param.ByName("productId")
	code := http.StatusNotModified
	if controller.Service.Delete(userId, productId) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}
