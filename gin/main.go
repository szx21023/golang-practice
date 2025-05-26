package main

import (
    "errors"
    "net/http"

    "github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Next()

        if len(c.Errors) > 0 {
            err := c.Errors.Last().Err

            c.JSON(http.StatusInternalServerError, map[string]any{
                "success": false,
                "message": err.Error(),
            })
        }
    }
}

func main() {
    r := gin.Default()
    r.Use(ErrorHandler())

    r.GET("/ok", func(c *gin.Context) {
        somethingWentWrong := false
        if somethingWentWrong {
            c.Error(errors.New("something went wrong"))
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "Everythin is fine!",
        })
    })

    r.GET("/error", func(c *gin.Context) {
        somethingWentWrong := true
        if somethingWentWrong {
            c.Error(errors.New("something wnet wrong"))
            return
        }

        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "message": "Everythin is fine!",
        })
    })

    r.Run(":8080")
}