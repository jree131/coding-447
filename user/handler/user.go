package handler

import (
	"context"
	"git.imocc.com/cap1573/user/domain/model"
	"git.imocc.com/cap1573/user/domain/service"
	user "git.imocc.com/cap1573/user/proto/user"
)

type User struct{

	UserDataService  service.IUserDataService

}

// 注册
func(u *User)Register(ctx context.Context, userRegiterRequest *user.UserRegiterRequest,userRegiterReponse *user.UserRegiterReponse) error{
	userRegister := &model.User{
		UserName: userRegiterRequest.UserName,
		FirstName: userRegiterRequest.FirstName,
		HashPassword: userRegiterRequest.Pwd,
	}
	_, err  := u.UserDataService.AddUser(userRegister)
	if  err != nil{
		return err
	}
	userRegiterReponse.Message="添加成功"
	return  nil

}
// 登录
func (u *User)Login(ctx context.Context,userLogin *user.UserLoginRequest,loginResponse *user.UserLoginReponse) error{
	isOk,err := u.UserDataService.CheckPwd(userLogin.UserName,userLogin.Pwd)
	if err !=nil {
		return err
	}
	loginResponse.IsSuccess = isOk
	return nil
}
// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context,userInfoRequest *user.UserInfoRequest, userInfoResponse *user.UserInfoReponse)  error{
     userInfo , err  := u.UserDataService.FindUserByName(userInfoRequest.UserName)

     if  err  != nil{
     	return  err
	 }

	 userInfoResponse  = UserForResponse(userInfo)
	 return nil
}
//类型转化
func UserForResponse(userModel *model.User) *user.UserInfoReponse  {
	response := &user.UserInfoReponse{}
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.ID
	return response
}
