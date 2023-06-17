package dto

type AddTransactionDTO struct {
	Data []*TransactionItem `json:"data"`
}

type TransactionItem struct {
	ProductId string `json:"productId"`
	Total     uint   `json:"total"`
}
