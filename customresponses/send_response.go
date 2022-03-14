package customresponses

import (
	"fmt"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data interface{}, code int) {
	response := NewByteResponseTemplate(data, code)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}
