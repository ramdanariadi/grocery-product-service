package ErrorHandlers

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/ramdanariadi/grocery-be-golang/main/customresponses"
	"github.com/ramdanariadi/grocery-be-golang/main/exception"
	"net/http"
)

func PanicHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			fmt.Printf("recover %v \n", err)

			if errors, ok := err.(validator.ValidationErrors); ok {
				var errMsg []string
				for _, err := range errors {
					errMsg = append(errMsg, err.Field()+" field is "+err.Tag())
				}
				customresponses.SendResponse(writer, errMsg, http.StatusBadRequest)
				return
			}

			if errors, ok := err.(exception.NotFoundError); ok {
				customresponses.SendResponse(writer, errors.Error(), http.StatusBadRequest)
				return
			}

			if errors, ok := err.(exception.NotFoundError); ok {
				customresponses.SendResponse(writer, errors.Error, http.StatusBadRequest)
				return
			}

			if err != nil {
				if errors, ok := err.(error); ok {
					customresponses.SendResponse(writer, errors.Error(), http.StatusInternalServerError)
					return
				}
				customresponses.SendResponse(writer, err, http.StatusInternalServerError)
				return
			}
		}()
		handler.ServeHTTP(writer, request)
	})
}

// for github.com/julienschmidt/httprouter v1.3.0 router.panicHandler
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
