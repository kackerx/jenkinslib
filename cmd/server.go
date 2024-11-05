package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	log.Println("start 8888")
	if err := http.ListenAndServe(":8888", r); err != nil {
		panic(err)
	}
}
