package history

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func init() {
	api.Register("POST", "history/search", historySearch)
}

type SearchBody struct {
	Tags   interface{}       `json:"tags,omitempty"`   //tags过滤器
	Begin  time.Time         `json:"begin"`            //开始时间
	End    time.Time         `json:"end"`              //结束时间
	Unit   string            `json:"unit,omitempty"`   //单位 year month week day hour minute second
	Step   int               `json:"step,omitempty"`   //步长
	Values map[string]string `json:"values,omitempty"` //values显示 sum avg min max first last median
}

func historySearch(ctx *gin.Context) {
	var body SearchBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//拼接查询流水
	var pipeline mongo.Pipeline
	filter := bson.D{{"tags", body.Tags}}
	if !body.Begin.IsZero() {
		filter = append(filter, bson.E{Key: "date", Value: bson.D{{"$gte", body.Begin}}})
	}
	if !body.End.IsZero() {
		filter = append(filter, bson.E{Key: "date", Value: bson.D{{"$lte", body.End}}})
	}
	match := bson.D{{"$match", filter}}
	pipeline = append(pipeline, match)

	//聚合
	groups := bson.D{{"_id", bson.D{{"$dateTrunc", bson.M{
		"date":        "$date",
		"unit":        body.Unit,
		"binSize":     body.Step,
		"timezone":    "+08:00", //time.Local.String(), Asia/Shanghai
		"startOfWeek": "monday",
		//TODO 改为系统时区
	}}}}}

	//取值
	for k, v := range body.Values {
		groups = append(groups, bson.E{
			Key:   k,
			Value: bson.D{{"$" + v, "$" + k}},
		})
	}

	pipeline = append(pipeline, bson.D{{"$group", groups}})

	//todo _id 重命名为 date

	var results []table.Document
	err = db.Aggregate(Bucket, pipeline, &results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}
