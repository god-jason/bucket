package web

import (
	"errors"
	"net/http"
	"path"
	"strings"
)

var Static FileSystem

type fileItem struct {
	fs    http.FileSystem
	path  string
	base  string
	index string
}

type FileSystem struct {
	items []*fileItem
	//items map[string]*fileItem
}

func (f *FileSystem) Put(path string, fs http.FileSystem, base string, index string) {
	f.items = append(f.items, &fileItem{fs: fs, path: path, base: base, index: index})
}

func (f *FileSystem) Open(name string) (file http.File, err error) {
	//低效
	for _, ff := range f.items {
		//fn := path.Join(ff.base, name)
		if ff.path == "" && !strings.HasPrefix(name, "/$") ||
			ff.path != "" && strings.HasPrefix(name, ff.path) {

			//去除前缀
			fn := path.Join(ff.base, strings.TrimPrefix(name, ff.path))

			//查找文件
			file, err = ff.fs.Open(fn)
			if file != nil {
				fi, _ := file.Stat()
				if !fi.IsDir() {
					return
				}
			}

			//尝试默认页
			if ff.index != "" {
				file, err = ff.fs.Open(path.Join(ff.base, ff.index))
				if file != nil {
					fi, _ := file.Stat()
					if !fi.IsDir() {
						return
					}
				}
			}

			return nil, errors.New("not found")
		}
	}
	return nil, errors.New("not found")
}
