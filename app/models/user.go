package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
	"time"
)

var UserClass = class.Make[User]()

func UserQuery() *table.Table[User] {
	return table.Class(UserClass, "users")
}

type User struct {
	Id        string    `json:"id"`
	NickName  string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetId 实现 auth 需要的方法
func (u User) GetId() string {
	return u.Id
}
