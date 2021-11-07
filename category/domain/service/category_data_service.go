package service

import (
	"git.imocc.com/cap1573/category/domain/model"
	"git.imocc.com/cap1573/category/domain/repository"
)

type ICategoryDataService  interface{
	AddCategory (*model.Category)  (int64  , error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64)  (*model.Category, error)
	FindCategoryByName(string) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)

}

// 创建
func  NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService{
	return  &CategoryDataService{categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}
// 插入
func (u  *CategoryDataService) AddCategory(category *model.Category)  (int64 ,error){
	return  u.CategoryRepository.CreateCategory(category)
}
// 删除
func  (u  *CategoryDataService) DeleteCategory(categoryID  int64)  error{
	return  u.CategoryRepository.DeleteCategory(categoryID)
}
// 更新
func (u *CategoryDataService) UpdateCategory(category  *model.Category) error{
	 return  u.CategoryRepository.UpdateCategory(category)
}
// 查找
func  (u *CategoryDataService) FindAllCategory() ([]model.Category, error){
	return u.CategoryRepository.FidAllCategory()
}
// 根据 ID　　查找
func  (u  *CategoryDataService) FindCategoryByID(CategoryID int64) (*model.Category,error){
	return  u.CategoryRepository.FindCategoryByID(CategoryID)
}

//根据名字查找
func (u  *CategoryDataService) FindCategoryByName(CategoryName string) (*model.Category,error){
	return  u.CategoryRepository.FindCategoryByName(CategoryName)
}
// 根据Level查找
 func  (u *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category,error){
 	return  u.CategoryRepository.FindCategoryByLevel(level)
 }
 // 根据Parent查找
func (u *CategoryDataService) FindCategoryByParent(parent int64)([]model.Category,error) {
	return u.CategoryRepository.FindCategoryByParent(parent)
}



