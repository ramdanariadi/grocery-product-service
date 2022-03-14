package category

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	FindById(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindAll(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}
