package wishlist

import (
	"net/http"
)

type CartController interface {
	FindByUserId(w http.ResponseWriter, r *http.Request)
	FindByUserAndProductId(w http.ResponseWriter, r *http.Request)
	Save(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
