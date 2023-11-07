package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayhello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
}
func main() {
	r := gin.Default() //返回默认的路由引擎
	//指定用户使用GET请求访问/hello时，执行sayhello这个函数
	r.GET("/hello", sayhello)
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mess": "GET",
		})
	})
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mess": "PSOT",
		})
	})
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mess": "PUT",
		})
	})
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"mess": "DELETE",
		})
	})
	//启动服务
	r.Run("127.0.0.1:9090")
}
