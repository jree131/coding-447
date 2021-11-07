package main

import (
	"fmt"
	"git.imocc.com/cap1573/user/domain/repository"
	service2 "git.imocc.com/cap1573/user/domain/service"
	"git.imocc.com/cap1573/user/handler"
	user "git.imocc.com/cap1573/user/proto/user"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-micro.dev/v4"
)

func main() {
	// 服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
	)

	// 初始化服务
	srv.Init()

	// 创建链接数据库
	db,err :=gorm.Open("mysql","root:123456@/micro?charset=utf8&parseTime=True&loc=Local")
	if  err != nil{
		fmt.Println(err)
	}

	defer db.Close()

	db.SingularTable(true)

	// 只执行一次数据库初始化
	//rp:= repository.NewUserRepository(db)
	//rp.InitTable()
	//创建服务实例
	userDataService := service2.NewUserDataService(repository.NewUserRepository(db))

	//注册Handler

	err = user.RegisterUserHandler(srv.Server(),&handler.User{UserDataService:userDataService})

	if  err  != nil{
		fmt.Println(err)
	}
	// 启动服务
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
