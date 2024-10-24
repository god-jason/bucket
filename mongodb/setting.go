package mongodb

import (
	"github.com/god-jason/bucket/setting"
	"github.com/god-jason/bucket/types"
)

func init() {
	setting.Register(MODULE, &setting.Module{
		Name:   "数据库",
		Module: MODULE,
		Title:  "数据库配置",
		Form: []types.SmartField{
			{Key: "url", Label: "连接字符串", Type: "text"},
			{Key: "database", Label: "数据库", Type: "text"},
			{Key: "auth", Label: "鉴权数据库", Type: "text"},
			{Key: "username", Label: "用户名", Type: "text"},
			{Key: "password", Label: "密码", Type: "text"},
		},
	})
}
