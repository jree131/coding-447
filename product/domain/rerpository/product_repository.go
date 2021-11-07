package rerpository

import (
	"git.imocc.com/cap1573/product/domain/model"
	"github.com/jinzhu/gorm"
)

type IProductRepository interface {
	InitTable() error
	FindProductByID(int64) (*model.Product, error)
	CreateProduct(*model.Product) (int64, error)
	UpdateProduct(product *model.Product) error
	FindAll() ([]model.Product, error)
	DeleteProductByID(int64) error
}

// 创建 productRepository

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u ProductRepository) InitTable() error {
	// 连续创建 4张表
	return u.mysqlDb.CreateTable(&model.Product{}, &model.ProductSeo{}, &model.ProductImage{}, &model.ProductSize{}).Error
}

// 根据ID　　查找 Product
func (u *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	//多表联查
	return product, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, productID).Error
}

//创建Product  信息
func (u ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, u.mysqlDb.Create(product).Error
}

func (u ProductRepository) UpdateProduct(product *model.Product) error {

	return u.mysqlDb.Model(product).Update(product).Error

}

// 根据ID　删除信息  开启事务
func (u ProductRepository) DeleteProductByID(productID int64) error {
	// 开启事务
	tx := u.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback() // 事务回滚
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}

	// 删除
	if err := tx.Unscoped().Where("id =?", productID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("images_product_id = ?", productID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("size_product_id = ?", productID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Unscoped().Where("seo_product_id = ?", productID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error

}

//获取结果集 多表查询
func (u ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, u.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productAll).Error
}
