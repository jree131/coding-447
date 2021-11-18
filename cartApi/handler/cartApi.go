package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cart "github.com/jree131/coding-447/cart/proto"
	cartApi "github.com/jree131/coding-447/cartApi/proto"
	"go-micro.dev/v4/util/log"
	"strconv"
)

type CartApi struct {
	CartService cart.CartService
}

// CartApi.Call 通过API向外暴露为/cartApi/findAll，接收http请求
// 即：/cartApi/call请求会调用go.micro.api.cartApi 服务的CartApi.Call方法

func (e *CartApi) FindAll(ctx context.Context, req *cartApi.Request, rep *cartApi.Response) error {
	log.Info("接受到 /cartApi/findAll 访问请求")
	//req.Get get请求
	if _, ok := req.Get["user_id"]; !ok {
		//rsp.StatusCode= 500
		return errors.New("参数异常")
	}

	userIDString := req.Get["user_id"].Values[0]
	fmt.Println(userIDString)
	userId, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		return err
	}
	// 获取购物车所有的商品
	cartAll, err := e.CartService.GetAll(context.TODO(), &cart.CartFindAll{
		UserId: userId,
	})
	b, err := json.Marshal(cartAll)
	if err != nil {
		return err
	}
	rep.StatusCode = 200
	rep.Body = string(b)
	return nil

}
