package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"time"
)

func GetCpuPercent() float64 {
	var percent, _ = cpu.Percent(time.Second, false) //使用量
	var count, _ = cpu.Info()                        //Mhz总频率
	Log.Info(count)
	Log.Info(percent)
	GetMemPercent()
	GetDiskPercent()
	return percent[0]
}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	Log.Info(memInfo)
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	Log.Info(parts)
	Log.Info(diskInfo)
	return diskInfo.UsedPercent
}
