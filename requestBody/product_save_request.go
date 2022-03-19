package requestBody

type ProductSaveRequest struct {
	Price       int64
	Weight      uint
	CategoryId  string
	PerUnit     int
	Description string
	ImageUrl    string
	Name        string
}
