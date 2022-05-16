package main

import (
	"net/http"

	"github.com/Sam-run/GoModule-GinDemo/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	// //Default返回一个默认的路由引擎
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	//
	// 	//输出结果
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run()
	router := gin.Default()
	//StaticFS，是加载完整的目录资源
	router.StaticFS("/public", http.Dir("/Users/shenzhu/Downloads/Go/GinDemo/web/static"))
	//StaticFile,是加载单个文件
	router.StaticFile("/logo.jpg", "./web/static/kaiyuan10nian.jpeg")

	//导入所有模版，多级目录结构
	router.LoadHTMLGlob("web/templete/*")

	//website分组
	v := router.Group("/")
	{
		v.GET("/index.html", controller.IndexController)
	}

	//http://localhost:8080/index.html

	router.Run()
}
