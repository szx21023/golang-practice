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
        userGroup.POST("/login", loginHandler)
    }
}

// 路由處理函數
func getUserHandler(c *gin.Context) {
    id := c.Param("id")
	// 回傳 id 給 client
    result, err := GetUserByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }
	c.JSON(http.StatusOK, gin.H{result.ID: result})
}

// 處理建立使用者的請求
func createUserHandler(c *gin.Context) {
    var req CreateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }
    result := CreateUser(req)

    // 使用結構化資料
    c.JSON(200, gin.H{"received": result})
}

func loginHandler(c *gin.Context) {
    var req LoginRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }

    // 模擬登入邏輯
    user, err := GetUserByAccount(req.Account)
    if err != nil || user.Password != req.Password {
        c.JSON(401, gin.H{"error": "Invalid credentials"})
        return
    }
    // 登入成功
    c.JSON(200, gin.H{"message": "Login successful", "user": user})
}