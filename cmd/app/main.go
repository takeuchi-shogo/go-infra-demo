package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
    A int
	B string
}

type Peple struct {}

func router() *gin.Engine {
		_ = Person{}
	p := Person{ A: 1, B: "taro" }
	log.Printf("%+v", p)
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello!!",
		})
	})
	           return router
}

func main() {
	if err := router().Run(); err != nil {
		panic(err)
	}
}
