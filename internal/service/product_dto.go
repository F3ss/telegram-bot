package service

var ListOfProduct []ProductDto = []ProductDto{
	{"One"},
	{"Two"},
	{"Three"},
	{"Four"},
	{"Five"},
}

type ProductDto struct {
	title string
}

func NewProductDto() *ProductDto {
	return &ProductDto{}
}
