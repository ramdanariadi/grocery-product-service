package customresponses

import (
	"encoding/json"
	"github.com/ramdanariadi/grocery-be-golang/main/helpers"
	"net/http"
)

type ResponseTemplate struct {
	Response interface{}     `json:"response"`
	MetaData MetaDataMessage `json:"metaData"`
}

func NewResponseTemplate(response interface{}, code int) *ResponseTemplate {
	metaData := MetaDataMessage{
		Code:    code,
		Message: http.StatusText(code),
	}
	return &ResponseTemplate{Response: response, MetaData: metaData}
}

func NewByteResponseTemplate(response interface{}, code int) []byte {
	cresponse := NewResponseTemplate(response, code)
	bytes, err := json.Marshal(&cresponse)
	helpers.PanicIfError(err)
	return bytes
}

type MetaDataMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
