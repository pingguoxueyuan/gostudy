package common

type UserInfo struct {
	UserId   uint64 `json:"user_id" db:"user_id"`
	Nickname string `json:"nickname" db:"nickname"`
	Sex      int    `json:"sex" db:"sex"`
	Username string `json:"user" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
