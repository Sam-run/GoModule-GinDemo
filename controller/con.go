package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Title": "开源十年",
	})
}
