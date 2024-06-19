package project

import (
	"github.com/god-jason/bucket/pkg/condition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AlarmScene struct {
	Id        primitive.ObjectID  `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID  `json:"project_id" bson:"project_id"`
	Name      string              `json:"name"`
	Times     []TimeRange         `json:"times"`
	Condition condition.Condition `json:"condition"`
	Alarm     Alarm               `json:"alarm"`
	Disabled  bool                `json:"disabled"`
}

type Alarm struct {
	Level   int    `json:"level,omitempty"`   //等级 1 2 3
	Type    string `json:"type,omitempty"`    //类型： 遥测 遥信 等
	Title   string `json:"title,omitempty"`   //标题
	Message string `json:"message,omitempty"` //内容
	//Template string `json:"template,omitempty"` //消息模板，format string
}
