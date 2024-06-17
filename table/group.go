package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	api.Register("POST", "table/:table/group", apiGroup)
}

type Group struct {
	Operator string `json:"operator,omitempty"` //sum count avg min max first last push
	Field    string `json:"field,omitempty"`
	As       string `json:"as,omitempty"`
}

type GroupBody struct {
	Filter map[string]interface{} `json:"filter,omitempty"`
	Field  string                 `json:"field,omitempty"`
	Format string                 `json:"format,omitempty"` //日期格式 支持 $dateToString %Y-%m-%d %H:%M:%S
	Groups []Group                `json:"groups,omitempty"`
}

func apiGroup(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	var body GroupBody
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//拼接查询流水
	var pipeline mongo.Pipeline

	match := bson.D{{"$match", body.Filter}}
	pipeline = append(pipeline, match)

	groups := bson.D{{"_id", "$" + body.Field}}
	for _, f := range table.Fields {
		if f.Name == body.Field {
			if f.Type == "date" {
				format := body.Format
				if format == "" {
					format = "%Y-%m-%d %H"
				}
				//日期类型要特殊处理
				groups = bson.D{{"_id", bson.D{{"$dateToString", bson.M{
					"format":   format,
					"date":     "$" + body.Field,
					"timezone": "+08:00", //time.Local.String(), Asia/Shanghai
					//TODO 改为系统时区
				}}}}}
			}
			break
		}
	}

	for _, g := range body.Groups {
		if g.Operator == "count" {
			groups = append(groups, bson.E{Key: g.As, Value: bson.D{{"$sum", 1}}})
		} else {
			groups = append(groups, bson.E{Key: g.As, Value: bson.D{{"$" + g.Operator, "$" + g.Field}}})
		}
	}
	group := bson.D{{"$group", groups}}
	pipeline = append(pipeline, group)

	var results []Document
	err = table.Aggregate(pipeline, &results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}
