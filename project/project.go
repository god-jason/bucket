package project

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name"`
	Disabled bool               `json:"disabled"`

	//定时任务
	timers []*TimerScene

	//条件任务 todo 监听设备
	conditions []*ConditionScene
}
