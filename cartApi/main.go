package main

import (
	"cartApi/handler"
	cartApi "cartApi/proto"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	opentracing2 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	cart "github.com/jree131/coding-447/cart/proto"
	"github.com/jree131/common"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"net"
	"net/http"
)

func main() {
	//注册中心
	consul := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.api.carApi", "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	// 启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", ""), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	//  new Service
	service := micro.NewService(
		micro.Name("go.mocro.api.cartApi"),
		micro.Version("latest"),
		micro.Address("0.0.0.0:8086"),
		// 添加 consul  注册中心
		micro.Registry(consul),
		// 添加链路追踪
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		//添加熔断
		micro.WrapClient(NewClientHystrixWrapper()),
		// 添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	service.Init()
	cartService := cart.NewCartService("go.micro.service.cart", service.Client())
	// Register Handler
	if err := cartApi.RegisterCartApiHandler(service.Server(), &handler.CartApi{CartService: cartService}); err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

// 添加熔断
func NewClientHystrixWrapper() client.Wrapper {

	return func(c client.Client) client.Client {
		return &clientWrapper{c}
	}

}
