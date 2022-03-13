package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

/*
参考：https://www.cnblogs.com/shuiche/p/9334225.html
 */
func main() {

	/*
	一、get请求
	1. 普通请求。
	2. 带参数请求
	 */
	r := gin.Default()
	r.GET("/test", func(content *gin.Context) {
		content.JSON(200, gin.H{
			"message": "test-1",
		})
	})
	r.GET("/t1", func(c *gin.Context) {
		c.String(200, "haha-t1")
	})

	/*
	1.1 query参数,url后面的参数；
		http://localhost:9090/t2?name=haha
		http://localhost:9090/t2?name=haha&age=28
	 */
	r.GET("/t2", func(c *gin.Context) {
		n := c.Query("name")
		m := c.DefaultQuery("age", "18")
		result := fmt.Sprintf("%s -- %s", n, m)
		c.String(200, result)
	})

	/*
	1.2 url参数
		http://localhost:9090/t3/es1 ，其中es1是n的值
	 */
	r.GET("t3/:key", func(c *gin.Context) {
		n := c.Param("key")
		c.String(200, n)
	})

	/*
	二、post请求
		curl -XPOST http://localhost:9090/p1 -H "Content-Type:application/x-www-form-urlencoded"  -d 'message="sdf"'
	 */
	r.POST("/p1", func(c *gin.Context) {
		message := c.PostForm("message")
		//c.String(200, message)
		c.JSON(200, gin.H{
			"result": message,
		})
	})

	r.Run("0.0.0.0:9090")

}