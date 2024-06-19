package project

import (
	"github.com/god-jason/bucket/pkg/condition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TimerScene 定时场景
type TimerScene struct {
	Id        primitive.ObjectID  `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID  `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID  `json:"space_id" bson:"space_id"`
	Name      string              `json:"name"`
	Timers    []Timer             `json:"timers"`
	Condition condition.Condition `json:"condition"` //组合条件
	Actions   []Action            `json:"actions"`   //动作
	Disabled  bool                `json:"disabled"`
}

// ConditionScene 联动场景
type ConditionScene struct {
	Id        primitive.ObjectID  `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID  `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID  `json:"space_id" bson:"space_id"`
	Name      string              `json:"name"`
	Times     []TimeRange         `json:"times,omitempty"`
	Condition condition.Condition `json:"condition"` //组合条件
	Actions   []Action            `json:"actions"`   //动作
	Disabled  bool                `json:"disabled"`
}

type Action struct {
	Devices    []string          `json:"devices"` //
	Action     string            `json:"action"`
	Parameters map[string]string `json:"parameters"`
}

type TimeRange struct {
	Start   int   `json:"start"`   //起始时间 每天的分钟
	End     int   `json:"end"`     //结束时间 每天的分钟
	Weekday []int `json:"weekday"` //0 1 2 3 4 5 6
}

type Timer struct {
	Clock   int   `json:"clock"`   //启动时间 每天的分钟 1440
	Weekday []int `json:"weekday"` //0 1 2 3 4 5 6
}
