package function

import (
	"encoding/json"
	"github.com/god-jason/bucket/lib"
	"github.com/god-jason/bucket/pkg/exception"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

var functions lib.Map[Function]

func Load(name string) error {

	fn := filepath.Join(viper.GetString("data"), Path, name+".json")
	buf, err := os.ReadFile(fn)
	if err != nil {
		return exception.Wrap(err)
	}

	var function Function
	err = json.Unmarshal(buf, &function)
	if err != nil {
		return exception.Wrap(err)
	}

	err = function.Compile()
	if err != nil {
		return exception.Wrap(err)
	}

	functions.Store(name, &function)

	return nil
}

func LoadAll() error {

	d := filepath.Join(viper.GetString("data"), Path)
	_ = os.MkdirAll(d, os.ModePerm)

	es, err := os.ReadDir(d)
	if err != nil {
		return exception.Wrap(err)
	}

	for _, e := range es {
		if e.IsDir() {
			continue
		}
		if filepath.Ext(e.Name()) == ".json" {
			err = Load(strings.TrimRight(e.Name(), ".json"))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
