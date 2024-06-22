package alarm

import (
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/log"
	"github.com/god-jason/bucket/pkg/errors"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Context struct {
	start int64 //发生时间
	times int   //重复次数
}

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

	contexts lib.Map[Context] //解绑设备，会导致内存泄露，虽然不大 todo 监听解绑
}

func (v *Validator) Open() error {
	if !v.DeviceId.IsZero() {
		dev := device.Get(v.DeviceId.Hex())
		if dev == nil {
			return errors.New("找不到设备")
		}
		dev.WatchValues(v)
	} else if !v.ProductId.IsZero() {
		if !v.SpaceId.IsZero() {
			spc := space.Get(v.SpaceId.Hex())
			if spc == nil {
				return errors.New("找不到空间")
			}
			spc.WatchValues(v)
		} else if !v.ProjectId.IsZero() {
			prj := project.Get(v.ProjectId.Hex())
			if prj == nil {
				return errors.New("找不到项目")
			}
			prj.WatchValues(v)
		} else {
			return errors.New("不能仅指定产品")
		}
	} else {
		return errors.New("没有指定设备和产品")
	}

	return v.Condition.Init()
}

func (v *Validator) Close() error {
	v.contexts.Clear()

	if !v.DeviceId.IsZero() {
		dev := device.Get(v.DeviceId.Hex())
		if dev != nil {
			dev.UnWatchValues(v)
		}
	}

	if !v.SpaceId.IsZero() {
		spc := space.Get(v.SpaceId.Hex())
		if spc != nil {
			spc.UnWatchValues(v)
		}
	}

	if !v.ProjectId.IsZero() {
		prj := project.Get(v.ProjectId.Hex())
		if prj != nil {
			prj.UnWatchValues(v)
		}
	}

	v.contexts.Clear()

	return nil
}

func (v *Validator) OnValuesChange(product, device primitive.ObjectID, values map[string]any) {
	if v.DeviceId.IsZero() {
		if v.ProductId.Hex() != product.Hex() {
			//不是当前产品
			return
		}
	} else {
		if v.DeviceId.Hex() != device.Hex() {
			//不是当前设备
			return
		}
	}

	ret, err := v.Condition.Eval(values)
	if err != nil {
		log.Error(err)
		return
	}

	//取上下文件
	ctx := v.contexts.Load(device.Hex())
	if ctx == nil {
		ctx = &Context{}
		v.contexts.Store(product.Hex(), ctx)
	}

	//条件为 假，则重置
	if !ret {
		ctx.start = 0
		ctx.times = 0
		return
	}

	//起始时间
	now := time.Now().Unix()
	if ctx.start == 0 {
		ctx.start = now
	}

	//延迟报警
	if v.Delay > 0 {
		if now < ctx.start+v.Delay {
			return
		}
	}

	if ctx.times > 0 {
		//重复报警
		if !v.Repeat {
			return
		}

		//超过最大次数
		if v.RepeatTimes > 0 && ctx.times >= v.RepeatTimes {
			return
		}

		//还没到时间
		if now < ctx.start+v.RepeatTimeout {
			return
		}

		ctx.start = now
	}
	ctx.times++

	//产生报警
	alarm := &Alarm{
		ProjectId: v.ProjectId,
		SpaceId:   v.SpaceId,
		ProductId: product, //v.ProductId,
		DeviceId:  device,  //v.DeviceId,支持同产品
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
