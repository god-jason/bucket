package user

import (
	"github.com/god-jason/bucket/db"
)

func init() {
	db.Register(new(User), new(Password))
}
