package handler

import (
	"context"
	"github.com/jree131/coding-447/cart/domain/model"
	"github.com/jree131/coding-447/cart/domain/service"
	cart "github.com/jree131/coding-447/cart/proto"
	"github.com/jree131/common"
)

type Cart struct {
	CartDataService service.ICartDataService
}

// 添加购物车
func (h *Cart) AddCart(ctx context.Context, request *cart.CartInfo, response *cart.ResponseAdd) (err error) {
	cart := &model.Cart{}
	common.SwapTo(request, cart)
	response.CartId, err = h.CartDataService.AddCart(cart)
	return err
}

// 清空购物车
func (h *Cart) CleanCart(ctx context.Context, request *cart.Clean, response *cart.Response) error {
	if err := h.CartDataService.CleanCart(request.User_Id); err != nil {
		return err
	}
	response.Msg = "购物车清空成功"
	return nil
}
func (h *Cart) Incr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.IncrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}

	response.Msg = "购物车添加成功"
	return nil
}

//购物车减少商品数量
func (h *Cart) Decr(ctx context.Context, request *cart.Item, response *cart.Response) error {
	if err := h.CartDataService.DecrNum(request.Id, request.ChangeNum); err != nil {
		return err
	}
	response.Msg = "购物程减少成功"
	return nil
}

//删除购物车
func (h *Cart) DeleteById(ctx context.Context, request *cart.CartID, response *cart.Response) error {
	if err := h.CartDataService.DeleteCart(request.Id); err != nil {
		return err
	}
	response.Msg = "购物车删除成功"
	return nil
}

//查询用户所有的购物车信息
func (h *Cart) GetAll(ctx context.Context, request *cart.CartFindAll, response *cart.CartAll) error {
	cartAll, err := h.CartDataService.FindAllCart(request.UserId)
	if err != nil {
		return err
	}

	for _, v := range cartAll {
		cart := &cart.CartInfo{}
		if err := common.SwapTo(v, cart); err != nil {
			return err
		}
		response.CartInfo = append(response.CartInfo, cart)
	}
	return nil
}
