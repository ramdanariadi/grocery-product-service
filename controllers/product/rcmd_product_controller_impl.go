package product

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-tunas/customresponses"
	"go-tunas/models"
	productrepositories "go-tunas/repositories/product"
	"go-tunas/services/product"
	"net/http"
)

type RcmdProductControllerImpl struct {
	Service product.RcmdProductService
}

func NewRcmdProductControllerImpl(db *sql.DB) *RcmdProductControllerImpl {
	return &RcmdProductControllerImpl{
		Service: product.RcmdProductServiceImpl{
			RcmdRepository: productrepositories.RcmdProductRepositoryImpl{
				DB: db},
			ProductRepository: productrepositories.ProductRepositoryImpl{DB: db},
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
