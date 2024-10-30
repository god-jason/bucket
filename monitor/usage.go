package monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/god-jason/bucket/api"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/mem"
)

func init() {
	api.Register("GET", "/usage/cpu", cpuStats)
	api.Register("GET", "/usage/memory", memStats)
	api.Register("GET", "/usage/disk", diskStats)
}

func memStats(ctx *gin.Context) {
	stat, err := mem.VirtualMemory()
	if err != nil {
		api.Error(ctx, err)
		return
	}
	api.OK(ctx, stat)
}

func cpuStats(ctx *gin.Context) {
	times, err := cpu.Times(false)
	if err != nil {
		api.Error(ctx, err)
		return
	}
	if len(times) == 0 {
		api.Fail(ctx, "查询失败")
		return
	}
	api.OK(ctx, times[0])
}

func diskStats(ctx *gin.Context) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		api.Error(ctx, err)
		return
	}

	usages := make([]*disk.UsageStat, 0)
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			api.Error(ctx, err)
			return
		}
		usages = append(usages, usage)
	}
	api.OK(ctx, usages)
}
