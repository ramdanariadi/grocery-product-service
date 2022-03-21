package requestBody

type TransactionSaveRequest struct {
	Id                string
	DetailTransaction []DetailTransaction
}

type DetailTransaction struct {
	ProductId string `json:"id"`
	Name      string `json:"name"`
	Wight     int    `json:"wight"`
	Price     int    `json:"price"`
	PerUnit   int    `json:"perUnit"`
	Total     int    `json:"total"`
}
