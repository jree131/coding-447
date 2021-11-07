
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get go-micro.dev/v4/cmd/protoc-gen-micro
	go get go-micro.dev/v4/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/user.proto

.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. user.proto
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild # 名称
dockerBuild:
	docker build . -t user:latest

