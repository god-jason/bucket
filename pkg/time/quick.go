package time

import (
	"sync"
	"time"
)

var quick Time
var once sync.Once

func Quick() Time {
	once.Do(func() {
		go func() {
			for {
				quick = Now()
				time.Sleep(time.Second)
			}
		}()
	})
	return quick
}
