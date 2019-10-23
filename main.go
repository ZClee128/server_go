package main

import "github.com/gin-gonic/gin"

func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		dict := make(map[string]string)
		dict["title"] = "你好"
		dict["image"] = "xxx.png"
		c.JSON(200, gin.H{
			"data" : dict,
			"message": "pong",
		})
	})
	r.POST("/Login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"code":  "0",
			"username": username,
			"password": password,
		})
	})
	r.Run() 
}