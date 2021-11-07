module git.imocc.com/cap1573/category

go 1.15

require (
	github.com/asim/go-micro/plugins/config/source/consul/v4 v4.0.0-20211028090348-ed690ed838cc
	github.com/asim/go-micro/plugins/registry/consul/v4 v4.0.0-20211028090348-ed690ed838cc
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/lib/pq v1.2.0 // indirect
	go-micro.dev/v4 v4.2.1
	golang.org/x/sys v0.0.0-20211025201205-69cdffdb9359 // indirect
	google.golang.org/protobuf v1.26.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
