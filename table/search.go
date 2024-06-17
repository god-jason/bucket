package table

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/curd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	api.Register("POST", "table/:table/search", apiSearch)
}

type SearchBody struct {
	Filter map[string]interface{} `json:"filter,omitempty"`
	Sort   map[string]int         `json:"sort,omitempty"`
	Skip   int64                  `json:"skip,omitempty"`
	Limit  int64                  `json:"limit,omitempty"`
	Fields map[string]int         `json:"fields,omitempty"`
	//Keyword string
}

func apiSearch(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	var body SearchBody
	err = ctx.ShouldBindJSON(&body)
	if err != nil {
		curd.Error(ctx, err)
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

	//寻找外键
	for _, f := range table.Fields {
		if body.Fields != nil {
			//没有查询的字段，不找外键
			if ff, ok := body.Fields[f.Name]; !ok || ff <= 0 {
				continue
			}
		}
		if f.Foreign != nil {
			lookup := bson.D{{"$lookup", bson.M{
				"from":         f.Foreign.Table,
				"localField":   f.Name,
				"foreignField": f.Foreign.Field,
				"as":           f.Foreign.As,
			}}}
			unwind := bson.D{{"$unwind", bson.M{
				"path": "$" + f.Foreign.As,
				//"includeArrayIndex":          "unwind_order_index",
				"preserveNullAndEmptyArrays": true,
			}}}
			pipeline = append(pipeline, lookup, unwind)

			if body.Fields != nil {
				body.Fields[f.Foreign.As] = 1
			}
		}
	}

	//显示字段
	if len(body.Fields) > 0 {
		project := bson.D{{"$project", body.Fields}}
		pipeline = append(pipeline, project)
	}

	var results []Document
	err = table.Aggregate(pipeline, &results)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, results)
}
