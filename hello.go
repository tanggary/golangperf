package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
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
	r.GET("/calc", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Result": calculate(3, 5),
		})
	})
	r.GET("/runup/:duration", func(c *gin.Context) {
		duration := c.Param("duration")
		int1, err := strconv.Atoi(duration)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		running(int1)
		c.JSON(http.StatusOK, gin.H{
			"Action": "Revving",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func calculate(x int, y int) int {
	return x + y
}

func running(x int) {
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

	time.Sleep(time.Second * time.Duration(x))
	close(done)
}
