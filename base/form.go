package base

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func FormFiles(ctx *gin.Context) (files []*multipart.FileHeader, err error) {
	form, err := ctx.MultipartForm()
	if err != nil {
		return nil, err
	}
	for _, f := range form.File {
		files = append(files, f...)
	}
	return
}
