package entity

import (
	"emcs-relay-go/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"strconv"
	"time"
)

type Hardware struct {
	Name       string `json:"name"`
	Total      string `json:"total"`
	Used       string `json:"used"`
	Proportion string `json:"proportion"`
}

func CPU() Hardware {
	var percent, _ = cpu.Percent(time.Second, false) //使用量
	var info, _ = cpu.Info()                         //Mhz总频率
	mhz := info[0].Mhz
	total := strconv.FormatInt(int64(mhz), 10)
	proportion := strconv.FormatInt(int64(percent[0]), 10)
	used := strconv.Itoa(int(info[0].Cores))
	return Hardware{Name: "CPU", Total: total, Used: used, Proportion: proportion}
}
func Memory() Hardware {
	info, _ := mem.VirtualMemory()
	total := strconv.FormatInt(int64(info.Total), 10)
	proportion := strconv.FormatInt(int64(info.UsedPercent), 10)
	used := strconv.FormatInt(int64(info.Used), 10)
	return Hardware{Name: "Memory", Total: total, Used: used, Proportion: proportion}
}
func Disk() Hardware {
	parts, _ := disk.Partitions(true)
	info, _ := disk.Usage(parts[0].Mountpoint)
	utils.Log.Info(parts)
	utils.Log.Info(info)
	total := strconv.FormatInt(int64(info.Total), 10)
	proportion := strconv.FormatInt(int64(info.UsedPercent), 10)
	used := strconv.FormatInt(int64(info.Used), 10)
	//MD /1000 /1000
	return Hardware{Name: "Disk", Total: total, Used: used, Proportion: proportion}
}
func Store() Hardware {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	utils.Log.Info(parts)
	utils.Log.Info(diskInfo)
	info, _ := mem.VirtualMemory()
	total := strconv.FormatInt(int64(info.Total), 10)
	proportion := strconv.FormatInt(int64(info.UsedPercent), 10)
	used := strconv.FormatInt(int64(info.Used), 10)
	return Hardware{Name: "Disk", Total: total, Used: used, Proportion: proportion}
}
