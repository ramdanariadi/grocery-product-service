package customresponses

type ResponseTemplate struct {
	Response interface{}     `json:"response"`
	MetaData MetaDataMessage `json:"metaData"`
}

func NewResponseTemplate(response interface{}, code int, message string) *ResponseTemplate {
	metaData := MetaDataMessage{
		Code:    code,
		Message: message,
	}
	return &ResponseTemplate{Response: response, MetaData: metaData}
}

type MetaDataMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
