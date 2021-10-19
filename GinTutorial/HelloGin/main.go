package main

import "github.com/gin-gonic/gin"

/**
 * @Author: sm.gu
 * @Email: sm.gu@aftership.com
 * @Date: 2021/9/5 1:54 下午
 * @Desc:
 */

func main() {
	r := gin.Default()

	r.GET("/hello", func(context *gin.Context) {
		context.String(200, "%s", "Hello Gin")
	})

	err := r.Run(":8090")
	if err != nil {
		return
	}
}
