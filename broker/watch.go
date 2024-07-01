package broker

import (
	"encoding/json"
	"github.com/god-jason/bucket/device"
	"github.com/god-jason/bucket/project"
	"github.com/god-jason/bucket/space"
)

type deviceWatcher map[string]int

func (w *deviceWatcher) OnDeviceValuesChange(product string, device string, values map[string]any) {
	buf, err := json.Marshal(map[string]any{
		"product": product,
		"device":  device,
		"values":  values,
	})
	if err != nil {
		return
	}
	_ = server.Publish("device/"+device+"/values", buf, false, 0)
}

var _deviceWatcher = deviceWatcher(map[string]int{})

func watchDeviceValues(id string) {
	dev := device.Get(id)
	if dev != nil {
		//dev.Count++
		_deviceWatcher[id] += 1
		dev.WatchValues(&_deviceWatcher)
	}
}

func unWatchDeviceValues(id string) {
	dev := device.Get(id)
	if dev != nil {
		//dev.Count++
		_deviceWatcher[id] -= 1
		if _deviceWatcher[id] <= 0 {
			dev.UnWatchValues(&_deviceWatcher)
		}
	}
}

type projectWatcher map[string]int

func (w *projectWatcher) OnProjectValuesChange(project, product, device string, values map[string]any) {
	buf, err := json.Marshal(map[string]any{
		"project": project,
		"product": product,
		"device":  device,
		"values":  values,
	})
	if err != nil {
		return
	}
	_ = server.Publish("project/"+project+"/values", buf, false, 0)
}

var _projectWatcher = projectWatcher(map[string]int{})

func watchProjectValues(id string) {
	prj := project.Get(id)
	if prj != nil {
		//prj.Count++
		_projectWatcher[id] += 1
		prj.WatchValues(&_projectWatcher)
	}
}

func unWatchProjectValues(id string) {
	prj := project.Get(id)
	if prj != nil {
		//prj.Count++
		_projectWatcher[id] -= 1
		if _projectWatcher[id] <= 0 {
			prj.UnWatchValues(&_projectWatcher)
		}
	}
}

type spaceWatcher map[string]int

func (w *spaceWatcher) OnSpaceValuesChange(space, product, device string, values map[string]any) {
	buf, err := json.Marshal(map[string]any{
		"space":   space,
		"product": product,
		"device":  device,
		"values":  values,
	})
	if err != nil {
		return
	}
	_ = server.Publish("space/"+space+"/values", buf, false, 0)
}

var _spaceWatcher = spaceWatcher(map[string]int{})

func watchSpaceValues(id string) {
	prj := space.Get(id)
	if prj != nil {
		//prj.Count++
		_spaceWatcher[id] += 1
		prj.WatchValues(&_spaceWatcher)
	}
}

func unWatchSpaceValues(id string) {
	prj := space.Get(id)
	if prj != nil {
		//prj.Count++
		_spaceWatcher[id] -= 1
		if _spaceWatcher[id] <= 0 {
			prj.UnWatchValues(&_spaceWatcher)
		}
	}
}
