package repository

import (
	"git.imocc.com/cap1573/category/domain/model"
	"github.com/jinzhu/gorm"
)


type ICategoryRepository interface {
	InitTable() error
	// 增
	CreateCategory(*model.Category) (int64,error)
	// 改
	UpdateCategory(*model.Category)  error
	// 删
	DeleteCategory(int64)  error
	//查
	FindCategoryByName(string2 string) (*model.Category,error)
	FindCategoryByID(int64) (*model.Category,error)
	FindCategoryByLevel(uint32) ([]model.Category,error)
	FindCategoryByParent(int64) ([]model.Category,error)
	FidAllCategory()([]model.Category,error)
}

//创建categoryRepository
func NewCategoryRepository(db *gorm.DB) ICategoryRepository  {
	return &CategoryRepository{mysqlDb:db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u *CategoryRepository) InitTable() error{
	return  u.mysqlDb.CreateTable(&model.Category{}).Error
}
// 根据ID查找Category 信息
func  (u *CategoryRepository) FindCategoryByID(categoryID int64 )(category *model.Category,err error){
	category = &model.Category{}
	return  category ,u.mysqlDb.First(category,categoryID).Error
}

// 创建Category 信息
 func (u *CategoryRepository) CreateCategory(category *model.Category) (int64, error){
 	return  category.ID , u.mysqlDb.Create(category).Error
 }
 // 根据ID删除 Category 信息
 func  (u *CategoryRepository) DeleteCategory(categoryID int64) error{
	 return u.mysqlDb.Where("id = ?",categoryID).Delete(&model.Category{}).Error
 }
 // 更新Category信息
 func(u  *CategoryRepository) UpdateCategory(category *model.Category) error{
 	return  u.mysqlDb.Model(category).Update(category).Error
 }

 // 获取结果集
func (u *CategoryRepository) FidAllCategory()(categoryAll []model.Category,err error) {
	return categoryAll, u.mysqlDb.Find(&categoryAll).Error
}
// 根据分类名称进行查询
func (u *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = &model.Category{}
	return category,u.mysqlDb.Where("category_name = ?",categoryName).Find(category).Error
}
func (u *CategoryRepository) FindCategoryByLevel (level uint32) (categorySlice []model.Category,err error) {
	return categorySlice,u.mysqlDb.Where("category_level = ?",level).Find(categorySlice).Error
}

func (u *CategoryRepository) FindCategoryByParent(parent int64)(categorySlice []model.Category,err error) {
	return categorySlice,u.mysqlDb.Where("category_parent = ?",parent).Find(categorySlice).Error
}





