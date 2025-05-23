package main

import (
    "log"
    "time"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Person struct {
    Name string `form:"name"`
    Address string `form:"address"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
    router := gin.Default()
    router.GET("/testing", startPage)
    router.Run(":8080")
}

func startPage(c *gin.Context) {
    var person Person

    err := c.ShouldBind(&person)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    log.Println(person.Name)
    log.Println(person.Address)
    log.Println(person.Birthday)

    c.String(200, "success")
}