package user

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// 註冊 user 模組相關的路由
func RegisterRoutes(r *gin.Engine) {
    userGroup := r.Group("/user")
    {
        userGroup.GET("/:id", getUserHandler)
        userGroup.POST("", createUserHandler)
    }
}

// 路由處理函數
func getUserHandler(c *gin.Context) {
    id := c.Param("id")
	// 回傳 id 給 client
	c.JSON(http.StatusOK, gin.H{"id": id})
}

// 處理建立使用者的請求
func createUserHandler(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    // 使用結構化資料
    c.JSON(200, gin.H{"received": req})
}
