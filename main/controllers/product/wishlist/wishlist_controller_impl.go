package wishlist

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/main/repositories/transactions"
	product3 "github.com/ramdanariadi/grocery-be-golang/main/services/product"
	"net/http"
)

type WishlistControllerImpl struct {
	Service product3.WishlistService
}

func NewWishlistController(db *sql.DB) *WishlistControllerImpl {
	return &WishlistControllerImpl{
		Service: product3.WishlistServiceImpl{
			WishlistRepository: transactions.WishlistRepositoryImpl{DB: db},
			ProductRepository:  product.ProductRepositoryImpl{DB: db},
		},
	}
}

func (controller WishlistControllerImpl) FindByUserId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	carts := controller.Service.FindByUserId(id)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller WishlistControllerImpl) FindByUserAndProductId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	carts := controller.Service.FindByUserAndProductId(userId, productId)
	customresponses.SendResponse(w, carts, http.StatusOK)
}

func (controller WishlistControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	code := http.StatusInternalServerError
	if controller.Service.Save(userId, productId) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller WishlistControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["userId"]
	productId := vars["productId"]
	code := http.StatusNotModified
	if controller.Service.Delete(userId, productId) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}
