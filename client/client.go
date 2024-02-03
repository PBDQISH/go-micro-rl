package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-micro-rl/client/controller"
	"time"
)

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"*"},
		AllowMethods: []string{"OPTIONS", "GET", "POST", "POST", "DELETE", "PUT"},
		MaxAge:       1 * time.Hour,
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")
	router.Use(Cors())
	router.GET("/user/login", controller.LoginPage)
	router.GET("/user/register", controller.RegisterPage)
	router.POST("/api/v1.0/userregister", controller.RegisterUser)
	router.POST("/api/v1.0/userlogin", controller.LoginUser)
	router.Run(":8066")
}
