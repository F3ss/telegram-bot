package service

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetProductList() []ProductDto {
	return ListOfProduct
}
