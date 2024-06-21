package scene

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
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

func (s *Scene) Open() error {
	//todo 找设备，注册变化 watch
	for _, c := range s.Conditions {
		for _, cc := range c {
			dev := device.Get(cc.DeviceId)
			if dev == nil {
				return errors.New("设备找不到")
			}
			dev.Watch(s)
		}
	}

	return s.Condition.Init()
}

func (s *Scene) Close() error {
	s.last = false
	return nil
}

func (s *Scene) OnDeviceValuesChange(m map[string]any) {

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
		//执行
		if !s.SpaceId.IsZero() {
			spc := space.Get(s.SpaceId.Hex())
			if spc == nil {
				spc.Execute(s.Actions)
			}
		} else if !s.ProjectId.IsZero() {
			prj := project.Get(s.ProjectId.Hex())
			if prj == nil {
				prj.Execute(s.Actions)
			}
		}

	}
	s.last = ret
}
