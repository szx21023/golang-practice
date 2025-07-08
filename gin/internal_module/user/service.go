package user

import "fmt"

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
