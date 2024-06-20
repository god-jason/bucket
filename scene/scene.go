package scene

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Time struct {
	Start   int   `json:"start"`             //起始时间 每天的分钟
	End     int   `json:"end"`               //结束时间 每天的分钟
	Weekday []int `json:"weekday,omitempty"` //0 1 2 3 4 5 6
}

// Scene 联动场景
type Scene struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id" bson:"space_id"`
	Name      string             `json:"name"`
	Times     []*Time            `json:"times,omitempty"`

	Condition //组合条件

	Actions  []*base.Action `json:"actions"` //动作
	Disabled bool           `json:"disabled"`

	last bool //上一次判断结果
}

func (s *Scene) Init() error {
	return s.Condition.Init()
}

func (s *Scene) Check() {
	//检查时间
	if len(s.Times) > 0 {
		now := time.Now()
		minute := now.Hour()*60 + now.Minute()
		weekday := now.Weekday()
		has := false
		for _, t := range s.Times {
			if t.Start < t.End {
				if minute < t.Start || minute > t.End {
					continue
				}
			} else {
				if minute < t.Start && minute > t.End {
					continue
				}
			}

			if len(t.Weekday) > 0 {
				ww := false
				for _, wd := range t.Weekday {
					if int(weekday) == wd {
						ww = true
						break
					}
				}
				if !ww {
					continue
				}
			}
			has = true
		}
		if !has {
			//todo
			return
		}
	}

	//检查条件
	ret, err := s.Condition.Eval()
	if err != nil {
		log.Error(err)
		return
	}

	if ret && !s.last {
		//todo 执行

	}
	s.last = ret
}
