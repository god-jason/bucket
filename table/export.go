package table

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strconv"
	"time"
)

func ApiExport(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var body SearchBody
	err = ctx.ShouldBindJSON(&body)
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

	var results []mongodb.Document
	err = table.AggregateDocument(pipeline, &results)
	if err != nil {
		Error(ctx, err)
		return
	}

	buf, err := json.Marshal(results)
	if err != nil {
		Error(ctx, err)
		return
	}

	filename := table.Name + time.Now().Format("20060102150405") + ".json"
	// 设置响应头
	ctx.Status(http.StatusOK)
	ctx.Header("Content-Type", "application/json")
	//ctx.Header("Content-Type", "application/octet-stream")
	ctx.Header("Content-Disposition", "attachment; filename="+filename)
	ctx.Header("Content-Length", strconv.Itoa(len(buf)))
	_, _ = ctx.Writer.Write(buf)
	//ctx.Abort()
}
