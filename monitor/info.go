package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/god-jason/bucket/pkg/build"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/super-l/machine-code/machine"
	"runtime"
)

func init() {
	api.Register("GET", "/info", info)
	api.Register("GET", "/info/cpu", cpuInfo)
	api.Register("GET", "/info/machine", machineInfo)
}

func info(ctx *gin.Context) {
	api.OK(ctx, gin.H{
		"version": build.Version,
		"build":   build.Build,
		"git":     build.GitHash,
		"gin":     gin.Version,
		"runtime": runtime.Version(),
	})
}

func cpuInfo(ctx *gin.Context) {
	info, err := cpu.Info()
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(info) == 0 {
		api.Fail(ctx, "查询失败")
		return
	}
	api.OK(ctx, info[0])
}

func machineInfo(ctx *gin.Context) {
	info := machine.GetMachineData()
	api.OK(ctx, gin.H{
		"sn":   info.BoardSerialNumber,
		"mac":  info.LocalMacInfo,
		"uuid": info.PlatformUUID,
		"cpu":  info.CpuSerialNumber,
	})
}
