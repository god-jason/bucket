package alarm

import (
	"errors"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Validator struct {
	Id        primitive.ObjectID `json:"_id" bson:"_id"`
	ProjectId primitive.ObjectID `json:"project_id" bson:"project_id"`
	SpaceId   primitive.ObjectID `json:"space_id" bson:"space_id"`
	ProductId primitive.ObjectID `json:"product_id" bson:"product_id"`
	DeviceId  primitive.ObjectID `json:"device_id" bson:"device_id"`

	Condition //直接嵌入条件

	Name     string `json:"name"`
	Level    int    `json:"level,omitempty"`   //等级 1 2 3
	Type     string `json:"type,omitempty"`    //类型： 遥测 遥信 等
	Title    string `json:"title,omitempty"`   //标题
	Message  string `json:"message,omitempty"` //内容
	Disabled bool   `json:"disabled"`

	Delay         int64 `json:"delay,omitempty"`
	Repeat        bool  `json:"repeat,omitempty"`
	RepeatTimeout int64 `json:"repeat_timeout,omitempty" bson:"repeat_timeout,omitempty"`
	RepeatTimes   int   `json:"repeat_times,omitempty" bson:"repeat_times,omitempty"`

	//last  bool  //上一次计算结果
	start int64 //发生时间
	times int   //重复次数
}

func (v *Validator) Init() error {

	if v.ProjectId.IsZero() {
		return errors.New("必须指定项目")
	}

	//过滤掉项目禁用的情况
	prj := project.Get(v.ProjectId.Hex())
	if prj == nil {
		return errors.New("找不到项目")
	}

	if !v.DeviceId.IsZero() {
		dev := device.Get(v.DeviceId.Hex())
		if dev == nil {
			return errors.New("找不到设备")
		}
		dev.Watch(v)
	} else if !v.ProductId.IsZero() {
		if !v.SpaceId.IsZero() {
			spc := space.Get(v.SpaceId.Hex())
			if spc == nil {
				return errors.New("找不到空间")
			}
			for _, dev := range spc.Devices() {
				dev.Watch(v)
			}
		} else {
			for _, dev := range prj.Devices() {
				dev.Watch(v)
			}
		}

	} else {
		return errors.New("没有指定设备和产品")
	}

	return v.Condition.Init()
}

func (v *Validator) OnDeviceValuesChange(ctx map[string]any) {
	ret, err := v.Condition.Eval(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	if !ret {
		v.start = 0
		v.times = 0
		return
	}

	//起始时间
	now := time.Now().Unix()
	if v.start == 0 {
		v.start = now
	}

	//延迟报警
	if v.Delay > 0 {
		if now < v.start+v.Delay {
			return
		}
	}

	if v.times > 0 {
		//重复报警
		if !v.Repeat {
			return
		}

		//超过最大次数
		if v.RepeatTimes > 0 && v.times >= v.RepeatTimes {
			return
		}

		//还没到时间
		if now < v.start+v.RepeatTimeout {
			return
		}

		v.start = now
	}
	v.times++

	//产生报警
	alarm := &Alarm{
		ProjectId: v.ProjectId,
		SpaceId:   v.SpaceId,
		ProductId: v.ProductId,
		DeviceId:  v.DeviceId,
		Level:     v.Level,
		Type:      v.Type,
		Title:     v.Title,
		Message:   v.Message,
		Created:   time.Now(),
	}

	_, err = _alarmTable.Insert(alarm)
	if err != nil {
		log.Error(err)
		return
	}

	//todo 发送 mqtt

}
