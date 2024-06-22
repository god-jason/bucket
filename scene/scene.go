package scene

import (
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/pool"
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

	deviceContainer base.DeviceContainer
	last            bool //上一次判断结果
}

func (s *Scene) Open() error {
	if !s.SpaceId.IsZero() {
		spc := space.Get(s.SpaceId.Hex())
		if spc != nil {
			spc.WatchValues(s)
			s.deviceContainer = spc
		} else {
			return errors.New("找不到空间")
		}
	} else if !s.ProjectId.IsZero() {
		prj := project.Get(s.ProjectId.Hex())
		if prj != nil {
			prj.WatchValues(s)
			s.deviceContainer = prj
		} else {
			return errors.New("找不到项目")
		}
	} else {
		return errors.New("无效场景")
	}

	//找设备，注册变化 watch
	//for _, c := range s.Conditions {
	//	for _, cc := range c {
	//		dev := device.Get(cc.DeviceId)
	//		if dev == nil {
	//			return errors.New("设备找不到")
	//		}
	//		dev.WatchValues(s)
	//	}
	//}

	return s.Condition.Init()
}

func (s *Scene) Close() error {
	s.last = false

	if !s.SpaceId.IsZero() {
		spc := space.Get(s.SpaceId.Hex())
		if spc != nil {
			spc.UnWatchValues(s)
		}
	}
	if !s.ProjectId.IsZero() {
		prj := project.Get(s.ProjectId.Hex())
		if prj != nil {
			prj.UnWatchValues(s)
		}
	}

	//找设备，unwatch
	//for _, c := range s.Conditions {
	//	for _, cc := range c {
	//		dev := device.Get(cc.DeviceId)
	//		if dev != nil {
	//			dev.UnWatchValues(s)
	//		}
	//	}
	//}

	return nil
}

func (s *Scene) OnValuesChange(product, device primitive.ObjectID, values map[string]any) {
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
			//没有合适的时间段
			return
		}
	}

	//检查条件
	ret, err := s.Condition.Eval()
	if err != nil {
		log.Error(err)
		return
	}

	//执行接口
	if ret && !s.last {
		//执行
		s.ExecuteIgnoreError()
	}
	s.last = ret
}

func (s *Scene) execute(id string, name string, params map[string]any) {
	dev := device.Get(id)
	if dev != nil {
		_ = pool.Insert(func() {
			_, err := dev.Action(name, params)
			if err != nil {
				log.Error(err)
			}
		})
	}
}

func (s *Scene) ExecuteIgnoreError() {
	for _, a := range s.Actions {
		if !a.DeviceId.IsZero() {
			s.execute(a.DeviceId.Hex(), a.Name, a.Parameters)
		} else if !a.ProductId.IsZero() {
			if s.deviceContainer != nil {
				ids, err := s.deviceContainer.Devices(a.ProductId)
				if err != nil {
					log.Error(err)
					continue
				}
				for _, d := range ids {
					s.execute(d.Hex(), a.Name, a.Parameters)
				}
			} else {
				log.Error("需要指定产品ID")
				//error
			}
		} else {
			//error
			log.Error("无效的动作")
		}
	}
}

func (s *Scene) Execute() error {
	for _, a := range s.Actions {
		if !a.DeviceId.IsZero() {
			dev := device.Get(a.DeviceId.Hex())
			if dev != nil {
				_, err := dev.Action(a.Name, a.Parameters)
				if err != nil {
					return err
				}
			} else {
				return errors.New("设备找不到")
			}
		} else if !a.ProductId.IsZero() {
			if s.deviceContainer != nil {
				ids, err := s.deviceContainer.Devices(a.ProductId)
				if err != nil {
					return err
				}
				for _, d := range ids {
					dev := device.Get(d.Hex())
					if dev != nil {
						_, err := dev.Action(a.Name, a.Parameters)
						if err != nil {
							return err
						}
					} else {
						return errors.New("设备找不到")
					}
				}
			} else {
				return errors.New("需要指定产品ID")
			}
		} else {
			return errors.New("无效的动作")
		}
	}
	return nil
}
