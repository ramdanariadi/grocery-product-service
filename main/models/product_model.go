package models

type ProductModelCSV struct {
	Id          string      `json:"id"`
	Deleted     bool        `json:"deleted"`
	Price       int64       `json:"price"`
	Weight      uint        `json:"weight"`
	CategoryId  string      `json:"categoryId"`
	PerUnit     int         `json:"perUnit"`
	Description string      `json:"description"`
	ImageUrl    interface{} `json:"imageUrl"`
	Name        string      `json:"name"`
}

type ProductModel struct {
	Id          string      `json:"id"`
	Price       uint64      `json:"price"`
	Weight      uint32      `json:"weight"`
	Category    string      `json:"category"`
	PerUnit     int64       `json:"perUnit"`
	Description string      `json:"description"`
	ImageUrl    interface{} `json:"imageUrl"`
	Name        string      `json:"name"`
	CategoryId  string      `json:"categoryId"`
}
