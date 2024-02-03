package controller

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	pb "go-micro-rl/service"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
	"net/http"
	"time"
)

var service = "User"
var opts client.CallOption = func(o *client.CallOptions) {
	o.RequestTimeout = time.Second * 30
	o.DialTimeout = time.Second * 30
}
var regData struct {
	Mobile   string `json:"mobile"`
	PassWord string `json:"password"`
}

func RegisterUser(ctx *gin.Context) {

	ctx.ShouldBindJSON(&regData)

	if regData.Mobile == "" || regData.PassWord == "" {
		fmt.Println("错误输入")
		return
	}

	consulReg := consul.NewRegistry(registry.Addrs("127.0.0.1:8500"))
	consulService := micro.NewService(
		micro.Client(client.NewClient()),
		micro.Registry(consulReg),
	)
	consulService.Init()

	microClient := pb.NewUserRegisterService(service, consulService.Client())

	fmt.Println("打印注册界面值：")
	fmt.Println(regData)

	resp, err := microClient.Register(context.TODO(), &pb.RegReq{
		Mobile:   regData.Mobile,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("用户注册，找不到远程服务！", err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func LoginUser(ctx *gin.Context) {
	// 获取数据
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}

	ctx.ShouldBindJSON(&regData)

	// 指定 consul 服务发现
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"))
	consulService := micro.NewService(
		micro.Client(client.NewClient()),
		micro.Registry(consulReg),
	)
	//初始化
	consulService.Init()
	// 初始化客户端
	microClient := pb.NewUserLoginService(service, consulService.Client())
	fmt.Println("打印用户输入的账户名和密码：")
	fmt.Println(regData)
	// 调用远程函数
	resp, err := microClient.Login(context.TODO(), &pb.RegReq{
		Mobile:   regData.Mobile,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("用户登录, 找不到远程服务!", err)
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}
