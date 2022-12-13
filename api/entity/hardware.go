package entity

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"os"
	"path/filepath"
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
	//utils.Log.Info(parts)
	//utils.Log.Info(info)
	total := strconv.FormatInt(int64(info.Total), 10)
	proportion := strconv.FormatInt(int64(info.UsedPercent), 10)
	used := strconv.FormatInt(int64(info.Used), 10)
	//MD /1000 /1000
	return Hardware{Name: "Disk", Total: total, Used: used, Proportion: proportion}
}
func Logs() Hardware {
	var size int64
	err := filepath.Walk("logs", func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		size = 0
	}
	gb := int64(1 * 1024 * 1022 * 1024)
	total := strconv.FormatInt(gb, 10)
	proportion := float64(size) / float64(gb)
	used := strconv.FormatInt(size, 10)
	return Hardware{Name: "LOGS", Total: total, Used: used, Proportion: fmt.Sprintf("%f", proportion)}
}
