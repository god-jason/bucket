package user

import (
	"time"
)

// User 用户
type User struct {
	Id       string    `json:"_id" bson:"_id"`
	Name     string    `json:"name,omitempty"`
	Username string    `json:"username,omitempty"`
	Admin    bool      `json:"admin,omitempty"`
	Disabled bool      `json:"disabled,omitempty"`
	Created  time.Time `json:"created,omitempty" xorm:"created"`
}

// Password 密码
type Password struct {
	Id       string `json:"_id" bson:"_id"`
	Password string `json:"password"`
}

type Role struct {
	Id          string    `json:"id" xorm:"pk"`
	Name        string    `json:"name,omitempty"`        //名称
	Description string    `json:"description,omitempty"` //说明
	Privileges  []string  `json:"privileges,omitempty"`
	Disabled    bool      `json:"disabled,omitempty"`
	Created     time.Time `json:"created" xorm:"created"`
}
