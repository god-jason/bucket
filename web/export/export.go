package export

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/db"
	"github.com/god-jason/bucket/web/curd"
	"go.mongodb.org/mongo-driver/bson"
)

func ApiExport(table, filename string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//id := ctx.MustGet("id")
		ids := ctx.QueryArray("id")

		var datum []bson.M

		err := db.Find(table, bson.D{{Key: "_id", Value: bson.E{Key: "$in", Value: ids}}}, nil, 0, 0, &datum)
		if err != nil {
			curd.Error(ctx, err)
			return
		}

		//下载头
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.Header("Content-Disposition", "attachment; filename="+filename+".zip") // 用来指定下载下来的文件名
		ctx.Header("Content-Transfer-Encoding", "binary")

		writer := zip.NewWriter(ctx.Writer)

		for _, data := range datum {
			id := data["id"]
			fn := fmt.Sprintf("%v.json", id)
			f, err := writer.Create(fn)
			if err != nil {
				return
			}

			buf, _ := json.Marshal(data)
			_, err = f.Write(buf)
			if err != nil {
				return
			}
		}

		err = writer.Close()
		if err != nil {
			return
		}
	}
}
