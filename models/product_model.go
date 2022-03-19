package models

type ProductModelCSV struct {
	Id          string
	Deleted     bool
	Price       int64
	Weight      uint
	CategoryId  string
	PerUnit     int
	Description string
	ImageUrl    string
	Name        string
}

type ProductModel struct {
	Id          string
	Price       int64
	Weight      uint
	Category    string
	PerUnit     int
	Description string
	ImageUrl    string
	Name        string
}
