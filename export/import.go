package export

import (
	"archive/zip"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/mongodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
)

func ApiImport(table string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		formFile, err := ctx.FormFile("file")
		if err != nil {
			api.Error(ctx, err)
			return
		}

		file, err := formFile.Open()
		if err != nil {
			api.Error(ctx, err)
			return
		}
		defer file.Close()

		reader, err := zip.NewReader(file, formFile.Size)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		var idss []primitive.ObjectID
		//数据解析
		for _, file := range reader.File {
			if file.FileInfo().IsDir() {
				continue
			}

			reader, err := file.Open()
			buf, err := io.ReadAll(reader)
			if err != nil {
				api.Error(ctx, err)
				return
			}

			var data []any
			err = json.Unmarshal(buf, &data)
			if err != nil {
				api.Error(ctx, err)
				return
			}

			ids, err := mongodb.InsertMany(table, data)
			if err != nil {
				api.Error(ctx, err)
				return
			}

			idss = append(idss, ids...)
		}

		api.OK(ctx, idss)
	}
}
