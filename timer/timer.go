package timer

import (
	"fmt"
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/pool"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"strconv"
	"strings"
)

// Timer 定时场景
type Timer struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id" bson:"space_id"`
	Name      string             `json:"name"`
	Clock     int                `json:"clock"`   //启动时间 每天的分钟 1440
	Weekday   []int              `json:"weekday"` //0 1 2 3 4 5 6
	Actions   []*base.Action     `json:"actions"` //动作
	Disabled  bool               `json:"disabled"`

	deviceContainer base.DeviceContainer
	entry           cron.EntryID
}

func (s *Timer) Open() (err error) {

	if !s.SpaceId.IsZero() {
		spc := space.Get(s.SpaceId.Hex())
		if spc != nil {
			s.deviceContainer = spc
		} else {
			return errors.New("找不到空间")
		}
	} else if !s.ProjectId.IsZero() {
		prj := project.Get(s.ProjectId.Hex())
		if prj != nil {
			s.deviceContainer = prj
		} else {
			return errors.New("找不到项目")
		}
	} else {
		return errors.New("无效场景")
	}

	//星期处理
	w := "*"
	if len(s.Weekday) > 0 {
		var ws []string
		for _, day := range s.Weekday {
			if day >= 0 && day <= 7 {
				ws = append(ws, strconv.Itoa(day))
			} else {
				//error
			}
		}
		w = strings.Join(ws, ",")
	}

	//分 时 日 月 星期
	spec := fmt.Sprintf("%d %d * * %s", s.Clock%60, s.Clock/60, w)
	s.entry, err = _cron.AddFunc(spec, func() {
		//池化 避免拥堵
		_ = pool.Insert(s.ExecuteIgnoreError)
	})
	return
}

func (s *Timer) Close() (err error) {
	_cron.Remove(s.entry)
	return
}

func (s *Timer) execute(id string, name string, params map[string]any) {
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

func (s *Timer) ExecuteIgnoreError() {
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

func (s *Timer) Execute() error {
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
