package service

import "github.com/f3ss/telegram-bot/internal/app/models"

type ProductService struct{}

func NewProductService() *ProductService {
	return &ProductService{}
}

func (p *ProductService) GetProductList() []models.ProductDto {
	return models.ListOfProduct
}
