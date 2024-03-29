package repository

import (
	"git.imocc.com/cap1573/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	// 初始化数据库表
	InitTable() error
   // 根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	// 根据id查找用户信息
	FindUserById(int64) (*model.User ,error)
	//创建用户
	CreateUser(*model.User) (int64,error)
	// 根据用户id删除用户
	DeleteUserById(int64) error
	// 更新用户信息
	UpdateUser(*model.User) error
	// 查找用户信息
	FindAll() ([]model.User,error)

}

//创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository  {
	return &UserRepository{mysqlDb:db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

// 初始化表
func (u  *UserRepository) InitTable() error{
	return  u.mysqlDb.CreateTable(&model.User{}).Error
}
// 根据用户名称查找用户的信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error){
	user = &model.User{}
	return  user, u.mysqlDb.Where("usr_name =?",name).Find(user).Error
}
// 根据用户ID查找用户的信息
func(u  *UserRepository) FindUserById(userID int64) (user *model.User, err error){
	user = &model.User{}
	return  user,u.mysqlDb.First(user,userID).Error
}
//创建用户
func(u *UserRepository) CreateUser(user *model.User) (userID int64,err error){
	return  user.ID ,u.mysqlDb.Create(user).Error
}

// 根据用户ID  删除用户
func (u *UserRepository )DeleteUserById(userID int64) error{
	return  u.mysqlDb.Where("id=?",userID).Delete(&model.User{}).Error
}
// 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error{
	return u.mysqlDb.Model(user).Update(&user).Error
}
// 查找用户信息
func (u *UserRepository)FindAll() (userAll []model.User,err error){
	return  userAll ,u.mysqlDb.Find(&userAll).Error
}



