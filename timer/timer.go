package timer

import (
	"fmt"
	"github.com/god-jason/bucket/base"
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
	Actions   []base.Action      `json:"actions"` //动作
	Disabled  bool               `json:"disabled"`

	entry cron.EntryID
}

func (s *Timer) Open() (err error) {
	//星期处理
	w := "*"
	if len(s.Weekday) > 0 {
		var ws []string
		for _, day := range s.Weekday {
			ws = append(ws, strconv.Itoa(day))
		}
		w = strings.Join(ws, ",")
	}

	//分 时 日 月 星期
	spec := fmt.Sprintf("%d %d * * %s", s.Clock%60, s.Clock/60, w)
	s.entry, err = _cron.AddFunc(spec, s.tick) //todo 池化
	return
}

func (s *Timer) Close() (err error) {
	_cron.Remove(s.entry)
	return
}

func (s *Timer) tick() {
	if !s.SpaceId.IsZero() {
		spc := space.Get(s.SpaceId.Hex())
		if spc != nil {
			spc.Execute(s.Actions)
		}
	} else if !s.ProjectId.IsZero() {
		prj := project.Get(s.ProjectId.Hex())
		if prj != nil {
			prj.Execute(s.Actions)
		}
	}
}
