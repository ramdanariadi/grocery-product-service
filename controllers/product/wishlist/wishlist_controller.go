package wishlist

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CartController interface {
	FindByUserId(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	FindByUserAndProductId(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Save(w http.ResponseWriter, r *http.Request, param httprouter.Params)
	Delete(w http.ResponseWriter, r *http.Request, param httprouter.Params)
}
