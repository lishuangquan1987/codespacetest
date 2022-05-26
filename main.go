package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("this is a go test")
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"Code":200,
			"Message":"test successfull",
		})
	})
	r.Run(":8090")
}
