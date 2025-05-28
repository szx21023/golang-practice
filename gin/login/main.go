package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	
	r.GET("/login_page", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Login Page",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login successfully!",
		})
	})

	r.Run("0.0.0.0:8080") // listen and serve on
}