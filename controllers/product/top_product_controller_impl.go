package product

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/customresponses"
	productrepositories "github.com/ramdanariadi/grocery-be-golang/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/services/product"
	"net/http"
)

type TopProductControllerImpl struct {
	Service product.TopProductService
}

func NewTopProductControllerImpl(db *sql.DB) *TopProductControllerImpl {
	return &TopProductControllerImpl{
		Service: product.TopProductServiceImpl{
			TopProductRepository: productrepositories.TopProductRepositoryImpl{
				DB: db},
		},
	}
}

func (controller TopProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productModel := controller.Service.FindById(id)
	customresponses.SendResponse(w, productModel, http.StatusOK)
}

func (controller TopProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productModels := controller.Service.FindAll()
	customresponses.SendResponse(w, productModels, http.StatusOK)
}

func (controller TopProductControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	code := http.StatusInternalServerError
	if controller.Service.Save(vars["id"]) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller TopProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	code := http.StatusNotModified
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, "", code)
}
