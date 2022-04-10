package product

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/models"
	product2 "github.com/ramdanariadi/grocery-be-golang/main/repositories/product"
	product3 "github.com/ramdanariadi/grocery-be-golang/main/services/product"
	"net/http"
)

type RcmdProductControllerImpl struct {
	Service product3.RcmdProductService
}

func NewRcmdProductControllerImpl(db *sql.DB) *RcmdProductControllerImpl {
	return &RcmdProductControllerImpl{
		Service: product3.RcmdProductServiceImpl{
			RcmdRepository: product2.RcmdProductRepositoryImpl{
				DB: db},
			ProductRepository: product2.ProductRepositoryImpl{DB: db},
		},
	}
}

func (controller RcmdProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productModel := controller.Service.FindById(id)

	if productModel == (models.ProductModel{}) {
		customresponses.SendResponse(w, nil, http.StatusNoContent)
		return
	}

	customresponses.SendResponse(w, productModel, http.StatusOK)
}

func (controller RcmdProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productModels := controller.Service.FindAll()
	customresponses.SendResponse(w, productModels, http.StatusOK)
}

func (controller RcmdProductControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code := http.StatusInternalServerError
	if controller.Service.Save(vars["id"]) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller RcmdProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	code := http.StatusInternalServerError
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, "", code)
}
