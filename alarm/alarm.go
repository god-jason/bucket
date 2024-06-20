package alarm

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Alarm struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id" bson:"space_id"`
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id"`
	DeviceId  primitive.ObjectID `json:"device_id" bson:"device_id"`

	Level   int    `json:"level,omitempty"`   //等级 1 2 3
	Type    string `json:"type,omitempty"`    //类型： 遥测 遥信 等
	Title   string `json:"title,omitempty"`   //标题
	Message string `json:"message,omitempty"` //内容

	Created time.Time `json:"created"`
}
