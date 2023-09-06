package main

import (
	"github.com/gin-gonic/gin"
	"xiangxiang/jackaroo/router"
)

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		router.Hello()
	})
	r.Run(":8000")
}
