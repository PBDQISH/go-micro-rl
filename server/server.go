package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro-rl/server/handler"
	"go-micro-rl/server/model"
	pb "go-micro-rl/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"
)

var (
	service = "User"
	//servicel = "userLogin"
	version = "latest"
)

func main() {
	//初始化数据库连接
	model.DbInit()

	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))

	service := micro.NewService(
		micro.Address("127.0.0.1:8800"),
		micro.Name(service),
		micro.Registry(consulReg),
		micro.Version(version),
	)
	service.Init()

	pb.RegisterUserRegisterHandler(service.Server(), new(handler.User))
	pb.RegisterUserLoginHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
