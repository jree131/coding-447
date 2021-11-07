package service

import (
	"git.imocc.com/cap1573/product/domain/model"
	"git.imocc.com/cap1573/product/domain/rerpository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FIndAllProduct() ([]model.Product, error)
}

// 创建
func NewProductDataService(productRepository rerpository.IProductRepository) IProductDataService {
	return &ProductDataService{productRepository}
}

type ProductDataService struct {
	ProductRepository rerpository.IProductRepository
}

// 插入
func (p ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return p.ProductRepository.CreateProduct(product)
}

// 删除
func (p ProductDataService) DeleteProduct(productID int64) error {
	return p.ProductRepository.DeleteProductByID(productID)
}
// 更新
func (p ProductDataService) UpdateProduct(product *model.Product) error {
	return p.ProductRepository.UpdateProduct(product)

}

// 根据ID查找
func (p ProductDataService) FindProductByID(productID int64) (*model.Product, error) {
  return  p.ProductRepository.FindProductByID(productID)
}

// 查找
func (p ProductDataService) FIndAllProduct() ([]model.Product, error) {
	return  p.ProductRepository.FindAll()
}
