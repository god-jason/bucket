package db

import (
	"github.com/god-jason/bucket/pkg/smart"
	"github.com/god-jason/bucket/setting"
)

func init() {
	setting.Register(MODULE, &setting.Module{
		Name:   "数据库",
		Module: MODULE,
		Title:  "数据库配置",
		Form: []smart.FormItem{
			{
				Key: "Type", Label: "数据库类型", Type: "select", Default: "mysql",
				Options: []smart.FormSelectOption{
					{Label: "SQLite（内置）", Value: "sqlite"},
					{Label: "MySQL", Value: "mysql"},
					{Label: "Postgres SQL", Value: "postgres"},
					{Label: "MS SQL Server", Value: "sqlserver"},
					{Label: "Oracle", Value: "godror"},
				},
			},
			{Key: "url", Label: "连接字符串", Type: "text"},
			{Key: "debug", Label: "调试模式", Type: "switch"},
		},
	})
}
