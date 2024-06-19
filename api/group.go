package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/table"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Group(tab *table.Table, after func(doc []table.Document) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body table.GroupBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		//拼接查询流水
		var pipeline mongo.Pipeline

		match := bson.D{{"$match", body.Filter}}
		pipeline = append(pipeline, match)

		groups := bson.D{{"_id", "$" + body.Field}}
		for _, f := range tab.Fields {
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
		err = tab.Aggregate(pipeline, &results)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(results)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, results)
	}
}
