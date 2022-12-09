package utils

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"os"
	"path/filepath"
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

// getFileSize get file size by path(B)
func DirSizeB(path string) (int64, error) {
	var size int64
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}

// getFileSize get file size by path(B)
func getFileSize(path string) int64 {
	if !exists(path) {
		return 0
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return fileInfo.Size()
}

// exists Whether the path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
