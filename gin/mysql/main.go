package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}
}