package dto

type FindWishlistDTO struct {
	Search    *string `form:"search"`
	PageIndex int     `form:"pageIndex"`
	PageSize  int     `form:"pageSize"`
}
