package aggregate

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	api.Register("POST", "aggregate/group", aggregateGroup)
}

func aggregateGroup(ctx *gin.Context) {
	var body table.GroupBody
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	//拼接查询流水
	var pipeline mongo.Pipeline

	match := bson.D{{"$match", body.Filter}}
	pipeline = append(pipeline, match)

	groups := bson.D{{"_id", "$" + body.Field}}
	for _, f := range _table.Fields {
		if f.Name == body.Field {
			if f.Type == "date" {
				groups = bson.D{{"_id", bson.D{{"$dateTrunc", bson.M{
					"date":        "$" + body.Field,
					"unit":        body.Unit,
					"binSize":     body.Step,
					"timezone":    viper.GetString("timezone"),
					"startOfWeek": "monday",
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

	var results []table.Document
	err = db.Aggregate(Bucket, pipeline, &results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}
