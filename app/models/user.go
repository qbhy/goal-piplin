package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var UserClass = class.Make[User]()

func Users() *table.Table[User] {
	return table.Class(UserClass, "users")
}

type User struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Role      string `json:"role"`
	Password  string `json:"password,omitempty"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

const (
	UserRoleAdmin = "admin"
	UserRoleUser  = "user"
)

// GetId 实现 auth 需要的方法
func (u User) GetId() string {
	return u.Id
}
