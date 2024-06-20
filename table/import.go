package table

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/base"
	"github.com/god-jason/bucket/db"
	"io"
)

func ApiImport(ctx *gin.Context) {
	table, err := Get(ctx.Param("table"))
	if err != nil {
		Error(ctx, err)
		return
	}

	var doc []db.Document

	//支持文件上传
	if ctx.ContentType() == "multipart/form-data" {
		files, err := base.FormFiles(ctx)
		if err != nil {
			Error(ctx, err)
			return
		}

		if len(files) != 1 {
			Fail(ctx, "仅支持一个文件")
			return
		}

		file, err := files[0].Open()
		defer file.Close()

		buf, err := io.ReadAll(file)
		if err != nil {
			Error(ctx, err)
			return
		}

		err = json.Unmarshal(buf, &doc)
		if err != nil {
			Error(ctx, err)
			return
		}
	} else {
		err := ctx.ShouldBind(&doc)
		if err != nil {
			Error(ctx, err)
			return
		}
	}

	ids, err := table.ImportDocument(doc)
	if err != nil {
		Error(ctx, err)
		return
	}

	OK(ctx, ids)
}
