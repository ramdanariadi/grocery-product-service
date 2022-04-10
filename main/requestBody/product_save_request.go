package requestBody

type ProductSaveRequest struct {
	Price       int64
	Weight      uint
	Category    string
	PerUnit     int
	Description string
	ImageUrl    string
	Name        string
}

type TopProductSaveRequest struct {
	ProductId   string
	Price       int64
	Weight      uint
	Category    string
	PerUnit     int
	Description string
	ImageUrl    interface{}
	Name        string
}

type RcmdProductSaveRequest struct {
	ProductId   string
	Price       int64
	Weight      uint
	Category    string
	PerUnit     int
	Description string
	ImageUrl    interface{}
	Name        string
}
