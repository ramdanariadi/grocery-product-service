package requestBody

type CategorySaveRequest struct {
	Category string `validate:"required"`
	ImageUrl string
}
