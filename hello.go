package main

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/calc", func(d *gin.Context) {
		d.JSON(http.StatusOK, gin.H{
			"Result": calculate(3, 5),
		})
	})
	r.GET("/runup", func(d *gin.Context) {
		running()
		d.JSON(http.StatusOK, gin.H{
			"Action": "Revving",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func calculate(x int, y int) int {
	return x + y
}

func running() {
	done := make(chan int)

	for i := 0; i < runtime.NumCPU(); i++ {
		go func() {
			for {
				select {
				case <-done:
					return
				default:
				}
			}
		}()
	}

	time.Sleep(time.Second * 10)
	close(done)
}
