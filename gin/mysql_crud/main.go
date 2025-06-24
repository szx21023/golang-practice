package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        int
	Name      string
}


func main() {
	dsn := "testuser:testpass@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"


	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("連線資料庫失敗:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ping 資料庫失敗:", err)
	}
	fmt.Println("成功連線 MySQL !")

	// 建立 users 資料表
	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL
	);`

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("建立資料表失敗:", err)
	}

	fmt.Println("成功建立 users 資料表！")

	r := gin.Default()
	r.GET("/create_customer", func(c *gin.Context) {
		name := c.Query("name")
		// 插入一筆使用者資料
		stmt, err := db.Prepare("INSERT INTO users (name) VALUES (?)")
		if err != nil {
			log.Fatal("預備語句失敗:", err)
		}
		defer stmt.Close()

		res, err := stmt.Exec(name)
		if err != nil {
			log.Fatal("執行插入失敗:", err)
		}

		lastID, err := res.LastInsertId()
		if err != nil {
			log.Fatal("取得 ID 失敗:", err)
		}

		fmt.Printf("成功建立使用者，ID: %d\n", lastID)

		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/read_customer", func(c *gin.Context) {
		rows, err := db.Query("SELECT id, name FROM users")
		if err != nil {
			log.Fatal("查詢失敗:", err)
		}
		defer rows.Close()

		var users []User

		for rows.Next() {
			var u User
			err := rows.Scan(&u.ID, &u.Name)
			if err != nil {
				log.Fatal("解析資料失敗:", err)
			}
			users = append(users, u)
		}

		// 顯示資料
		for _, u := range users {
			fmt.Printf("ID: %d, Name: %s\n", u.ID, u.Name)
		}
	})
	r.Run("0.0.0.0:8080")
}