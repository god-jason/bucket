package api

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/table"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Search(tab *table.Table, after func(doc []db.Document) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var body table.SearchBody
		err := ctx.ShouldBindJSON(&body)
		if err != nil {
			Error(ctx, err)
			return
		}

		//拼接查询流水
		var pipeline mongo.Pipeline

		match := bson.D{{"$match", body.Filter}}
		pipeline = append(pipeline, match)

		if body.Sort != nil && len(body.Sort) > 0 {
			sort := bson.D{{"$sort", body.Sort}}
			pipeline = append(pipeline, sort)
		}
		if body.Skip > 0 {
			skip := bson.D{{"$skip", body.Skip}}
			pipeline = append(pipeline, skip)
		}
		if body.Limit > 0 {
			limit := bson.D{{"$limit", body.Limit}}
			pipeline = append(pipeline, limit)
		}

		var fields map[string]int
		if len(body.Fields) > 0 {
			for _, f := range body.Fields {
				fields[f] = 1
			}
		}
		//寻找外键
		for _, f := range tab.Fields {
			if fields != nil {
				//没有查询的字段，不找外键
				if _, ok := fields[f.Name]; !ok {
					continue
				}
			}
			if f.Foreign != nil {
				lookup := bson.D{{"$lookup", bson.M{
					"from":         f.Foreign.Table,
					"localField":   f.Name,
					"foreignField": f.Foreign.Field,
					"as":           f.Foreign.As,
					"pipeline": bson.D{
						{"$limit", 1},
						{"$project", bson.M{f.Foreign.Let: 1}},
					},
				}}}
				//数组扁平化
				unwind := bson.D{{"$unwind", bson.M{
					"path": "$" + f.Foreign.As,
					//"includeArrayIndex":          "unwind_order_index",
					"preserveNullAndEmptyArrays": true, //保留空，实现左外连接
				}}}
				//提高一级
				set := bson.D{{"$set", bson.M{f.Foreign.As: "$" + f.Foreign.As + "." + f.Foreign.Let}}}

				pipeline = append(pipeline, lookup, unwind, set)

				if fields != nil {
					fields[f.Foreign.As] = 1
				}
			}
		}

		//显示字段
		if len(body.Fields) > 0 {
			project := bson.D{{"$project", fields}}
			pipeline = append(pipeline, project)
		}

		var results []db.Document
		err = tab.AggregateDocument(pipeline, &results)
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
