package user

// 模擬的資料模型
type User struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    Account string `json:"account"`
    Password string `json:"password"`
}