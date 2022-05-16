package main

import (
	"net/http"

	"github.com/Sam-run/GoModule-GinDemo/controller"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func main() {

	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	type user struct {
		gorm.Model
		account  string
		password string
		name     string
		mobile   uint
	}
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
