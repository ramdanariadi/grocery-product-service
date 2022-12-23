package model

type TransactionModel struct {
	Id                string                           `json:"id"`
	UserId            string                           `json:"user_id"`
	Name              string                           `json:"name"`
	TotalPrice        uint64                           `json:"total_price"`
	TransactionDate   string                           `json:"transaction_date"`
	DetailTransaction []*DetailTransactionProductModel `json:"detail_transaction"`
}

type DetailTransactionProductModel struct {
	Id            string `json:"id"`
	ProductId     string `json:"productId"`
	TransactionId string `json:"transactionId"`
	Price         uint64 `json:"price"`
	Weight        uint   `json:"weight"`
	Category      string `json:"repository"`
	PerUnit       uint   `json:"perUnit"`
	Description   string `json:"description"`
	ImageUrl      string `json:"imageUrl"`
	Name          string `json:"name"`
	CategoryId    string `json:"categoryId"`
	Total         uint   `json:"total"`
}
