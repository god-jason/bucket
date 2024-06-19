package timer

import (
	"github.com/god-jason/bucket/base"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Timer 定时场景
type Timer struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id" bson:"space_id"`
	Name      string             `json:"name"`
	Clock     int                `json:"clock"`   //启动时间 每天的分钟 1440
	Weekday   []int              `json:"weekday"` //0 1 2 3 4 5 6
	Actions   []base.Action      `json:"actions"` //动作
	Disabled  bool               `json:"disabled"`
}
