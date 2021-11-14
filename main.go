package main

// go get -u github.com/gin-gonic/gin
import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("💻 -- Running on localhost:8080 using Gin.")
	fmt.Println("💻 -- try http://localhost:8080/ping")

	// Default With the Logger and Recovery middleware already attached
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// add test framework
}
