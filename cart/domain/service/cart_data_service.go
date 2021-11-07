package service

import (
	"github.com/jree131/coding-447/cart/domain/model"
	"github.com/jree131/coding-447/cart/domain/repository"
)

type ICartDataService interface {
	AddCart(*model.Cart) (int64, error)
	DeleteCart(int64) error
	UpdateCart(*model.Cart) error
	FindCartByID(int64) (*model.Cart, error)
	FindAllCart(int64) ([]model.Cart, error)

	CleanCart(int64) error
	DecrNum(int64, int64) error
	IncrNum(int64, int64) error
}

// 创建
func NewCartDataService(cartRepositoy repository.ICartRepository) ICartDataService {
	return &CartDataService{cartRepositoy}
}

type CartDataService struct {
	CartRepository repository.ICartRepository
}

// 插入
func (u *CartDataService) AddCart(cart *model.Cart) (int64, error) {
	return u.CartRepository.CreateCat(cart)
}

// 删除
func (u *CartDataService) DeleteCart(cartID int64) error {
	return u.CartRepository.DeleteCarByID(cartID)
}

// 更新
func (u *CartDataService) UpdateCart(cart *model.Cart) error {
	return u.CartRepository.UpdateCart(cart)
}

// 查找
func (u *CartDataService) FindCartByID(cartID int64) (*model.Cart, error) {
	return u.CartRepository.FindCartByID(cartID)
}

//  查找所有
func (u *CartDataService) FindAllCart(userID int64) ([]model.Cart, error) {
	return u.CartRepository.FindAll(userID)
}

func (u *CartDataService) CleanCart(userID int64) error {
	return u.CartRepository.CleanCart(userID)
}

func (u *CartDataService) DecrNum(cartID int64, num int64) error {
	return u.CartRepository.DecrNum(cartID, num)
}

func (u *CartDataService) IncrNum(cartID int64, num int64) error {
	return u.CartRepository.IncNum(cartID, num)
}
