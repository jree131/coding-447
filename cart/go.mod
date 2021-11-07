module github.com/jree131/coding-447/cart

go 1.15

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4 v4.0.0-20211103025805-c5be9f560cdb
	github.com/golang/protobuf v1.5.2
	github.com/jinzhu/gorm v1.9.16
	github.com/jree131/common v0.0.0-20211106084846-332daf6055a3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
