package user

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("your_secret_key") // JWT 密鑰
// 模擬的資料庫
var users = map[string]User{}

// 根據 ID 取得使用者
func GetUserByID(id string) (User, error) {
    user, exists := users[id]
    if !exists {
        return User{}, fmt.Errorf("user not found")
    }
    return user, nil
}

func GetUserByAccount(account string) (User, error) {
    for _, user := range users {
        if user.Account == account {
            return user, nil
        }
    }
    return User{}, fmt.Errorf("user not found")
}

// 建立使用者
func CreateUser(req CreateUserRequest) User {
    id := fmt.Sprintf("%d", len(users)+1)
    user := User{
        ID:   id,
        Name: req.Name,
        Account: req.Account,
        Password: req.Password,
    }
    users[id] = user
    return user
}

func GenerateToken(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}
