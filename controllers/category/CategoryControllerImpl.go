package category

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-tunas/customresponses"
	categoryCustomResponse "go-tunas/customresponses/category"
	categoryRepository "go-tunas/repositories/category"
	"go-tunas/services/category"
	"net/http"
)

type CategoryControllerImpl struct {
	service category.CategoryService
}

func NewCategoryController(db *sql.DB) CategoryController {
	return &CategoryControllerImpl{
		service: category.CategoryServiceImpl{
			Repository: categoryRepository.CategoryRepositoryImpl{
				DB: db,
			},
		},
	}
}

func (controller CategoryControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {

}

func (controller CategoryControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	categories := controller.service.FindAll()

	var categoryResponse []categoryCustomResponse.CategoryResponse
	for _, categoryItem := range categories {
		cresponseTmp := categoryCustomResponse.CategoryResponse{
			Id:       categoryItem.Id,
			Category: categoryItem.Category,
			ImageUrl: categoryItem.ImageUrl,
		}
		categoryResponse = append(categoryResponse, cresponseTmp)
	}

	response := customresponses.NewResponseTemplate(categoryResponse, 200, "OK")

	jsonResponse, err := json.Marshal(&response)

	if err != nil {
		panic("json.Marshall error")
	}

	fmt.Fprintf(w, string(jsonResponse))
}

func (controller CategoryControllerImpl) Save(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller CategoryControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func (controller CategoryControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}
