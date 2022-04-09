package product

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ramdanariadi/grocery-be-golang/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/helpers"
	"github.com/ramdanariadi/grocery-be-golang/models"
	productrepositories "github.com/ramdanariadi/grocery-be-golang/repositories/product"
	"github.com/ramdanariadi/grocery-be-golang/requestBody"
	"github.com/ramdanariadi/grocery-be-golang/services/product"
	"io"
	"net/http"
)

type ProductControllerImpl struct {
	Service product.ProductService
}

func NewProductControllerImpl(db *sql.DB) *ProductControllerImpl {
	return &ProductControllerImpl{
		Service: product.ProductServiceImpl{
			Repository: productrepositories.ProductRepositoryImpl{
				DB: db},
		},
	}
}

func (controller ProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productModel := controller.Service.FindById(id)

	if productModel == (models.ProductModel{}) {
		customresponses.SendResponse(w, nil, http.StatusNoContent)
		return
	}
	customresponses.SendResponse(w, productModel, http.StatusOK)
}

func (controller ProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productModels := controller.Service.FindAll()
	customresponses.SendResponse(w, productModels, http.StatusOK)
}

func (controller ProductControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.ProductSaveRequest{}
	err = json.Unmarshal(bytes, &productSaveRequest)
	helpers.PanicIfError(err)

	code := http.StatusInternalServerError
	if controller.Service.Save(productSaveRequest) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller ProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.ProductSaveRequest{}
	err = json.Unmarshal(bytes, &productSaveRequest)
	helpers.PanicIfError(err)

	vars := mux.Vars(r)
	id := vars["id"]

	code := http.StatusNotModified
	if controller.Service.Update(productSaveRequest, id) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}

func (controller ProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	code := http.StatusNotModified
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, "", code)
}
