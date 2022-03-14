package category

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"go-tunas/customresponses"
	categoryCustomResponse "go-tunas/customresponses/category"
	"go-tunas/exception"
	"go-tunas/helpers"
	categoryRepository "go-tunas/repositories/category"
	"go-tunas/requestBody"
	"go-tunas/services/category"
	"io"
	"net/http"
)

type CategoryControllerImpl struct {
	Service  category.CategoryService
	Validate *validator.Validate
}

func NewCategoryController(db *sql.DB) CategoryController {
	return &CategoryControllerImpl{
		Service: category.CategoryServiceImpl{
			Repository: categoryRepository.CategoryRepositoryImpl{
				DB: db,
			},
		},
		Validate: validator.New(),
	}
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	category := controller.Service.FindById(param.ByName("id"))

	cresponseTmp := categoryCustomResponse.CategoryResponse{
		Id:       category.Id,
		Category: category.Category,
		ImageUrl: category.ImageUrl,
	}

	customresponses.SendResponse(w, cresponseTmp, http.StatusOK)
}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	categories := controller.Service.FindAll()

	var categoryResponse []categoryCustomResponse.CategoryResponse
	for _, categoryItem := range categories {
		cresponseTmp := categoryCustomResponse.CategoryResponse{
			Id:       categoryItem.Id,
			Category: categoryItem.Category,
			ImageUrl: categoryItem.ImageUrl,
		}
		categoryResponse = append(categoryResponse, cresponseTmp)
	}
	customresponses.SendResponse(w, categoryResponse, http.StatusOK)
}

func (controller CategoryControllerImpl) Save(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	panic(exception.NewNotFoundError("tes"))
	body := r.Body
	byte, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	saveRequest := requestBody.CategorySaveRequest{}
	err = json.Unmarshal(byte, &saveRequest)
	helpers.PanicIfError(err)

	err = controller.Validate.Struct(&saveRequest)
	helpers.PanicIfError(err)

	code := http.StatusInternalServerError
	if controller.Service.Save(saveRequest) {
		code = http.StatusCreated
	}

	customresponses.SendResponse(w, nil, code)
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")

	body := r.Body
	byte, err := io.ReadAll(body)
	helpers.PanicIfError(err)

	saveRequest := requestBody.CategorySaveRequest{}
	json.Unmarshal(byte, &saveRequest)

	code := http.StatusInternalServerError
	if controller.Service.Update(saveRequest, id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, nil, code)
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	id := param.ByName("id")

	code := http.StatusNotModified

	if controller.Service.Delete(id) {
		code = http.StatusOK
	}

	customresponses.SendResponse(w, nil, code)
}
