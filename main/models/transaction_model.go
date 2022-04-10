package models

type TransactionModel struct {
	Id                string
	UserId            string
	Name              string
	Total             int64
	DetailTransaction []ProductModel
}
