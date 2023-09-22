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

	r.GET("/runup/:duration/:percentage", func(c *gin.Context) {
		duration := c.Param("duration")
		int1, err := strconv.Atoi(duration)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		percentage := c.Param("percentage")
		int2, err := strconv.Atoi(percentage)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		fmt.Println(strconv.Itoa(runtime.NumCPU()) + " " + duration + " " + percentage)
		runcpuload(runtime.NumCPU(), int1, int2)

		str_cat := "Revving" + " " + percentage + "% " + duration + " seconds"

		c.JSON(http.StatusOK, gin.H{
			"Action": str_cat,
		})
	})

	r.GET("/maxup/:duration", func(c *gin.Context) {
		duration := c.Param("duration")
		int1, err := strconv.Atoi(duration)
		if err != nil {
			fmt.Println("Error during conversion")
			return
		}
		maxup(int1)

		str_cat := "Revving" + " 100%" + " " + duration + " seconds"

		c.JSON(http.StatusOK, gin.H{
			"Action": str_cat,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func runcpuload(coresCount int, timeSeconds int, percentage int) {
	runtime.GOMAXPROCS(coresCount)

	// second     ,s  * 1
	// millisecond,ms * 1000
	// microsecond,μs * 1000 * 1000
	// nanosecond ,ns * 1000 * 1000 * 1000

	// every loop : run + sleep = 1 unit

	// 1 unit = 100 ms may be the best
	unitHundresOfMicrosecond := 1000
	runMicrosecond := unitHundresOfMicrosecond * percentage
	sleepMicrosecond := unitHundresOfMicrosecond*100 - runMicrosecond

	done := make(chan int)

	for i := 0; i < coresCount; i++ {
		go func() {
			runtime.LockOSThread()
			// endless loop
			for {
				select {
				case <-done:
					return
				default:
					begin := time.Now()
					for {
						// run 100%
						if time.Now().Sub(begin) > time.Duration(runMicrosecond)*time.Microsecond {
							break
						}
					}
					// sleep
					time.Sleep(time.Duration(sleepMicrosecond) * time.Microsecond)
				}
			}
		}()
	}
	// how long
	time.Sleep(time.Duration(timeSeconds) * time.Second)
	close(done)
}

func maxup(x int) {
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
