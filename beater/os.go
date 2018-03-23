package beater

import (
	pro "github.com/shirou/gopsutil/process"
	"encoding/json"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/cpu"
)

type (
	OSProcessInfo struct {
		Cmdline       string                   `json:"cmdline"`
		Connections   []net.ConnectionStat     `json:"connections"`
		NetIOCounters []net.IOCountersStat     `json:"net_io_counters"`
		Threads       map[int32]*cpu.TimesStat `json:"threads"`
		IOCounters    *pro.IOCountersStat      `json:"io_counters"`
		CPUPercent    float64                  `json:"cpu_percent"`
		CreateTime    int64                    `json:"create_time"`
		MemoryInfo    *pro.MemoryInfoStat      `json:"memory_info"`
		MemoryPercent float32                  `json:"memory_percent"`
		NumThreads    int32                    `json:"num_threads"`
	}
)

func GetOSProcessInfo(pid int32) (string, error) {
	process, err := pro.NewProcess(pid)
	if IsOSError(err) {
		return "NewProcess", err
	}
	processInfo := new(OSProcessInfo)

	processInfo.Cmdline, err = process.Cmdline()
	if IsOSError(err) {
		return "Cmdline", err
	}

	processInfo.Connections, err = process.Connections()
	if IsOSError(err) {
		return "Connections", err
	}

	// Not implemented on darwin, windows
	processInfo.NetIOCounters, err = process.NetIOCounters(true)
	if IsOSError(err) {
		return "NetIOCounters", err
	}
	// Not implemented on darwin, windows
	processInfo.Threads, err = process.Threads()
	if IsOSError(err) {
		return "Threads", err
	}

	// Not implemented on darwin, windows
	processInfo.IOCounters, err = process.IOCounters()
	if IsOSError(err) {
		return "IOCounters", err
	}

	processInfo.CPUPercent, err = process.CPUPercent()
	if IsOSError(err) {
		return "CPUPercent", err
	}

	processInfo.CreateTime, err = process.CreateTime()
	if IsOSError(err) {
		return "CreateTime", err
	}

	processInfo.MemoryInfo, err = process.MemoryInfo()
	if IsOSError(err) {
		return "MemoryInfo", err
	}

	processInfo.MemoryPercent, err = process.MemoryPercent()
	if IsOSError(err) {
		return "MemoryPercent", err
	}

	processInfo.NumThreads, err = process.NumThreads()
	if IsOSError(err) {
		return "NumThreads", err
	}

	result, err := json.Marshal(processInfo)
	if err != nil {
		return "json.Marshal", err
	}

	return string(result), nil
}
