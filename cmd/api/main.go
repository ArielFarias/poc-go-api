package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.New()

	router.GET("/", func(context *gin.Context) {
		context.String(200, "Hello World!")
	})

	router.Run(":8080")
}