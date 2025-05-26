package main

import (
    "log"
    "time"

    "github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        c.Set("example", "12345")

        c.Next()

        latency := time.Since(t)
        log.Print(latency)

        status := c.Writer.Status()
        log.Println(status)
    }
}

func main() {
    router := gin.New()
    router.Use(Logger())

    router.GET("/test", func(c *gin.Context) {
        example := c.MustGet("example").(string)

        log.Println(example)
    })

    router.Run(":8080")
}