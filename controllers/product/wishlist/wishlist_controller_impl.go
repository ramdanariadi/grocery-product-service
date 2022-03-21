package wishlist

import (
	"database/sql"
	"github.com/julienschmidt/httprouter"
	"go-tunas/customresponses"
	"go-tunas/repositories/product"
	"go-tunas/repositories/transactions"
	product2 "go-tunas/services/product"
	"net/http"
)

type WishlistControllerImpl struct {
	Service product2.WishlistService
}

func NewWishlistController(db *sql.DB) *WishlistControllerImpl {
	return &WishlistControllerImpl{
		Service: product2.WishlistServiceImpl{
			WishlistRepository: transactions.WishlistRepositoryImpl{DB: db},
			ProductRepository:  product.ProductRepositoryImpl{DB: db},
		},
	}
}

func (controller WishlistControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("userId")
	carts := controller.Service.FindByUserId(id)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller WishlistControllerImpl) FindByUserAndProductId(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	productId := param.ByName("productId")
	carts := controller.Service.FindByUserAndProductId(userId, productId)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller WishlistControllerImpl) Save(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	productId := param.ByName("productId")
	code := http.StatusInternalServerError
	if controller.Service.Save(userId, productId) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller WishlistControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	userId := param.ByName("userId")
	productId := param.ByName("productId")
	code := http.StatusNotModified
	if controller.Service.Delete(userId, productId) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}
