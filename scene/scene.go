package scene

import (
	"github.com/god-jason/bucket/action"
	"github.com/god-jason/bucket/pkg/condition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Scene 联动场景
type Scene struct {
	Id        primitive.ObjectID  `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID  `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID  `json:"space_id" bson:"space_id"`
	Name      string              `json:"name"`
	Times     []Time              `json:"times,omitempty"`
	Condition condition.Condition `json:"condition"` //组合条件
	Actions   []action.Action     `json:"actions"`   //动作
	Disabled  bool                `json:"disabled"`
}

type Time struct {
	Start   int   `json:"start"`             //起始时间 每天的分钟
	End     int   `json:"end"`               //结束时间 每天的分钟
	Weekday []int `json:"weekday,omitempty"` //0 1 2 3 4 5 6
}
