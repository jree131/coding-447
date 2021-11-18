module github.com/jree131/coding-447/cartApi

go 1.15

require (
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4 v4.0.0-20211108090337-8e312801a106
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/golang/protobuf v1.5.2
	github.com/jree131/coding-447/cart v0.0.0-20211107084300-46b8c1fa12eb
	github.com/jree131/common v0.0.0-20211106084846-332daf6055a3
	github.com/opentracing/opentracing-go v1.2.0
	go-micro.dev/v4 v4.2.1
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
