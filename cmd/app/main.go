package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello!!",
		})
	})
	return router
}

func main() {
	router().Run()
}
