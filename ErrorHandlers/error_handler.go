package ErrorHandlers

import (
	"github.com/go-playground/validator/v10"
	"go-tunas/customresponses"
	"go-tunas/exception"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	if errors, ok := err.(validator.ValidationErrors); ok {
		var errMsg []string
		for _, err := range errors {
			errMsg = append(errMsg, err.Field()+" field is "+err.Tag())
		}
		customresponses.SendResponse(w, errMsg, http.StatusBadRequest)
		return
	}

	if errors, ok := err.(exception.NotFoundError); ok {
		customresponses.SendResponse(w, errors.Error, http.StatusBadRequest)
		return
	}

	customresponses.SendResponse(w, err.(error).Error(), http.StatusInternalServerError)
}
