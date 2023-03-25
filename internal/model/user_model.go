package model

type User struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credintails struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
