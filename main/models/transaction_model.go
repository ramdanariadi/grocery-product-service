package models

import "google.golang.org/genproto/googleapis/type/date"

type TransactionModel struct {
	Id                string
	UserId            string
	Name              string
	TotalPrice        uint64
	TransactionDate   date.Date
	DetailTransaction []DetailTransactionProductModel
}

type DetailTransactionProductModel struct {
	Id            string      `json:"id"`
	ProductId     string      `json:"productId"`
	TransactionId string      `json:"transactionId"`
	Price         uint64      `json:"price"`
	Weight        uint32      `json:"weight"`
	Category      string      `json:"category"`
	PerUnit       uint64      `json:"perUnit"`
	Description   string      `json:"description"`
	ImageUrl      interface{} `json:"imageUrl"`
	Name          string      `json:"name"`
	CategoryId    string      `json:"categoryId"`
	Total         uint        `json:"total"`
}
