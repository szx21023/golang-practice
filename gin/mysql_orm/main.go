package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID uint `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("連線資料庫失敗")
	}

	db.AutoMigrate(&User{})

	r := gin.Default()
	r.GET("/create_customer", func(c *gin.Context) {
		name := c.Query("name")
		// 插入一筆使用者資料
		newUser := User{Name: name}
		result := db.Create(&newUser)
		if result.Error != nil {
			fmt.Println("插入資料失敗")
		}

		c.JSON(200, gin.H{
			"message": "成功新增使用者",
			"user":    newUser,
		})
	})
	r.GET("/get_customers", func(c *gin.Context) {
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			fmt.Println("查詢資料失敗")
		}

		c.JSON(200, gin.H{
			"message": "成功取得使用者列表",
			"users":   users,
		})
	})
	r.Run("0.0.0.0:8080")

}