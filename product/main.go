package main

import (
	"git.imocc.com/cap1573/product/common"
	"git.imocc.com/cap1573/product/domain/rerpository"
	"git.imocc.com/cap1573/product/domain/service"
	"git.imocc.com/cap1573/product/handler"
	product "git.imocc.com/cap1573/product/proto"
	consul2 "github.com/asim/go-micro/plugins/registry/consul/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/jinzhu/gorm"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go-micro.dev/v4/util/log"
)

func main() {
	// 配置中心
	consulConfig ,err := common.GetConsulConfig("127.0.0.1",8500 ,"micro/config")
	if  err  != nil{
		log.Error(err)
	}
	// 注册中心
  consul:= consul2.NewRegistry(func(options *registry.Options) {
  	options.Addrs = []string{
  		"127.0.0.1:8500",
	}
  })

  // 链路追踪
  t,io, err := common.NewTracer("go.micro.service.product","localhost:6831")
  if  err  !=nil{
  	log.Fatal(err)
  }
  defer  io.Close()
  opentracing.SetGlobalTracer(t)

  // 数据库设置
  mysqlInfo := common.GetMysqlFromConsul(consulConfig,"mysql")
	db,err := gorm.Open("mysql",mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+")/" +mysqlInfo.Database+"?charset=utf8&parseTime=True&loc=Local")
    if  err  != nil{
    	log.Error(err)
	}
	defer  db.Close()

	// 禁止副表
	db.SingularTable(true)
	//初始化
	rerpository.NewProductRepository(db).InitTable()
	productDataService :=service.NewProductDataService(rerpository.NewProductRepository(db))

	// 设置服务
	service := micro.NewService(
		  micro.Name("go.micro.service.product"),
		  micro.Version("latest"),
		  micro.Address("127.0.0.1:8005"),
		  // 添加注册中心
		  micro.Registry(consul),
		  	// 绑定链路追踪
		  	micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		)

	// Initialise service
	service.Init()

	// Register Handler
	product.RegisterProductHandler(service.Server(), &handler.Product{ProductDataService:productDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}


}

