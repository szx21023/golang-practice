package user

type CreateUserRequest struct {
    Name string `json:"name"`
    Account string `json:"account"`
    Password string `json:"password"`
}

type LoginRequest struct {
    Account string `json:"account"`
    Password string `json:"password"`
}