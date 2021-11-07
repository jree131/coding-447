package repository

import (
	"errors"
	"git.imooc.com/coding-447/cart/domain/model"
	"github.com/jinzhu/gorm"
)

type ICartRepository interface {
	InitTable() error
	FindCartByID(int64) (*model.Cart, error)
	CreateCat(*model.Cart) (int64, error)
	DeleteCarByID(int64) error
	UpdateCart(*model.Cart) error
	FindAll(int64) ([]model.Cart, error)

	CleanCart(int64) error
	IncNum(int64, int64) error
	DecrNum(int64, int64) error
}

// 创建 cartRepository
func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{mysqlDb: db}
}

type CartRepository struct {
	mysqlDb *gorm.DB
}


// 初始化表
func (u *CartRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.Cart{}).Error
}

// 根据id查找 Cart信息
func (u *CartRepository) FindCartByID(carID int64) (cart *model.Cart, err error) {
	cart = &model.Cart{}
	return cart, u.mysqlDb.First(cart, carID).Error
}

// 创建Cart 信息
func (u *CartRepository) CreateCat(cart *model.Cart) (int64,error) {
  db := u.mysqlDb.FirstOrCreate(cart,model.Cart{ProductID: cart.ProductID ,
  	 SizeID: cart.SizeID ,UserID: cart.UserID})
  if  db.Error !=nil{
  	return  0, db.Error
  }
  if  db.RowsAffected ==0{
  	return 0, errors.New("购物车插入失败！")
  }
  return  cart.ID ,nil
}

// 根据ID删除Cart信息
func (u *CartRepository) DeleteCarByID(cartID int64) error{
	return  u.mysqlDb.Where("id  =?",cartID).Delete(&model.Cart{}).Error
}

// 更新Cart信息
func (u *CartRepository) UpdateCart(cart *model.Cart) error{
	return  u.mysqlDb.Model(cart).Update(cart).Error
}

// 获取结果集
func (u *CartRepository) FindAll(userID int64) (cartAll []model.Cart,err  error){
	return  cartAll,u.mysqlDb.Where("user_id =?" , userID).Find(&cartAll).Error
}

// 根据用户ID清空购物车
func(u  *CartRepository) CleanCart(userID  int64) error{
	return  u.mysqlDb.Where("user_id",userID).Delete(&model.Cart{}).Error
}

// 添加商品数量
func (u *CartRepository) IncNum(cartID  int64 , num  int64)error{
 cart  :=model.Cart{ID :cartID}
 return  u.mysqlDb.Model(cart).UpdateColumn("num",gorm.Expr("num + ?",num)).Error
}

// 购物车减少商品
func (u *CartRepository) DecrNum(cartID int64 , num int64) error{
	cart := &model.Cart{ID : cartID}
	db := u.mysqlDb.Model(cart).Where("num >= ?" , num).UpdateColumn("num",gorm.Expr("num - ?",num))

	if  db.Error != nil{
		return db.Error
	}

	if  db.RowsAffected ==0{
		return errors.New("减少失败")
	}
	return nil
}




