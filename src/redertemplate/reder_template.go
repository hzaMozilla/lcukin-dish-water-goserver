package redertemplate

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RederTemplate() {
	var r = gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
	fmt.Println("测试")
}