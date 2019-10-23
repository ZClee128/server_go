package main

import "github.com/gin-gonic/gin"
import "fmt"
import "io/ioutil"
import "encoding/json"

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
		data, _ := ioutil.ReadAll(c.Request.Body)
		fmt.Printf("ctx.Request.body: %v", string(data))
		var d map[string]interface{}
		// 将字符串反解析为字典
		json.Unmarshal(data, &d)
		fmt.Println(d)   
		c.JSON(200, gin.H{
			"code":  "0",
			"data": d,
		})
	})
	r.Run() 
}