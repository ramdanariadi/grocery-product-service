package product

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	productrepositories "go-tunas/repositories/product"
	"go-tunas/requestBody"
	"go-tunas/services/product"
	"io"
	"net/http"
)

type RcmdProductControllerImpl struct {
	Service product.RcmdProductService
}

func NewRcmdProductControllerImpl(db *sql.DB) *RcmdProductControllerImpl {
	return &RcmdProductControllerImpl{
		Service: product.RcmdProductServiceImpl{
			Repository: productrepositories.RcmdProductRepositoryImpl{
				DB: db},
		},
	}
}

func (controller RcmdProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	productModel := controller.Service.FindById(id)
	customresponses.SendResponse(w, productModel, http.StatusOK)
}

func (controller RcmdProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	productModels := controller.Service.FindAll()
	customresponses.SendResponse(w, productModels, http.StatusOK)
}

func (controller RcmdProductControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.RcmdProductSaveRequest{}
	err = json.Unmarshal(bytes, &productSaveRequest)
	helpers.PanicIfError(err)

	code := http.StatusInternalServerError
	if controller.Service.Save(productSaveRequest) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller RcmdProductControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.RcmdProductSaveRequest{}
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

func (controller RcmdProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	code := http.StatusInternalServerError
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, "", code)
}
