package main

import (
    "github.com/gin-gonic/gin"
    "internal_module/user" // 假設 module 名為 project
)

func main() {
    r := gin.Default()

    // 註冊 user 模組路由
    user.RegisterRoutes(r)

    // 啟動伺服器
    r.Run(":8080")
}
