package service

import (
	"errors"
	"git.imocc.com/cap1573/user/domain/model"
	"git.imocc.com/cap1573/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(*model.User) (int64, error)

	DeleteUser(int64) error

	UpdateUser(user *model.User, isChangePwd bool) (err error)

	FindUserByName(string) (*model.User, error)

	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

// 创建实例
func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

// 加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

// 验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误！")
	}
	return true, nil
}

// 插入用户
func (u *UserDataService) AddUser(user *model.User) (userId int64, err error) {
	// 加密用户密码
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)

}

// 删除用户
func (u *UserDataService) DeleteUser(userID int64) error {
	return u.UserRepository.DeleteUserById(userID)
}

// 更新用户
func (u *UserDataService) UpdateUser(user *model.User, isChangePwd bool) (err error) {
	// 判断用户是否更新密码
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}

		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)

}

// 根据用户名称查找用户
func (u *UserDataService) FindUserByName(userName string) (user *model.User, err error) {
	return u.UserRepository.FindUserByName(userName)
}

// 比对账户密码是否正确
func (u *UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}

	return ValidatePassword(pwd, user.HashPassword)
}
