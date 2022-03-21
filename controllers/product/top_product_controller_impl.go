package product

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go-tunas/customresponses"
	"go-tunas/helpers"
	productrepositories "go-tunas/repositories/product"
	"go-tunas/requestBody"
	"go-tunas/services/product"
	"io"
	"net/http"
)

type TopProductControllerImpl struct {
	Service product.TopProductService
}

func NewTopProductControllerImpl(db *sql.DB) *TopProductControllerImpl {
	return &TopProductControllerImpl{
		Service: product.TopProductServiceImpl{
			Repository: productrepositories.TopProductRepositoryImpl{
				DB: db},
		},
	}
}

func (controller TopProductControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("id")
	productModel := controller.Service.FindById(id)
	customresponses.SendResponse(w, productModel, http.StatusOK)
}

func (controller TopProductControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	productModels := controller.Service.FindAll()
	customresponses.SendResponse(w, productModels, http.StatusOK)
}

func (controller TopProductControllerImpl) Save(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.TopProductSaveRequest{}
	err = json.Unmarshal(bytes, &productSaveRequest)
	helpers.PanicIfError(err)

	code := http.StatusInternalServerError
	if controller.Service.Save(productSaveRequest) {
		code = http.StatusCreated
	}
	customresponses.SendResponse(w, "", code)
}

func (controller TopProductControllerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	body := r.Body
	bytes, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	productSaveRequest := requestBody.TopProductSaveRequest{}
	err = json.Unmarshal(bytes, &productSaveRequest)
	helpers.PanicIfError(err)

	id := param.ByName("id")

	code := http.StatusNotModified
	if controller.Service.Update(productSaveRequest, id) {
		code = http.StatusOK
	}
	customresponses.SendResponse(w, "", code)
}

func (controller TopProductControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	id := param.ByName("id")

	code := http.StatusInternalServerError
	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, "", code)
}
