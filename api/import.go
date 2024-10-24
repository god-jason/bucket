package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/mongodb"
	"github.com/god-jason/bucket/table"
	"io"
)

func Import(tab *table.Table, after func(ids []string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var doc []mongodb.Document

		//支持文件上传
		if ctx.ContentType() == "multipart/form-data" {
			files, err := table.FormFiles(ctx)
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

		ids, err := tab.ImportDocument(doc)
		if err != nil {
			Error(ctx, err)
			return
		}

		if after != nil {
			err = after(ids)
			if err != nil {
				Error(ctx, err)
				return
			}
		}

		OK(ctx, ids)
	}
}
