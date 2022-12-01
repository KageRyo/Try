package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	//設定靜態資源的讀取
	server.Static("/assets", "./template/assets")
	server.GET("/login", LoginPage)
	server.Run(":8888")
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
