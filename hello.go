package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/runup", func(d *gin.Context) {
		d.JSON(http.StatusOK, gin.H{
			"Result": calculate(3, 5),
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func calculate(x int, y int) int {
	return x + y
}
