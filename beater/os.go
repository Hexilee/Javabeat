package beater

import (
	pro "github.com/shirou/gopsutil/process"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/cpu"
)

type (
	OSProcessInfo struct {
		Cmdline       string                   `json:"cmdline"`
		Status        string                   `json:"status"`
		Connections   []net.ConnectionStat     `json:"connections"`
		NetIOCounters []net.IOCountersStat     `json:"net_io_counters"`
		Threads       map[int32]*cpu.TimesStat `json:"threads"`
		Times         *cpu.TimesStat           `json:"times"`
		IOCounters    *pro.IOCountersStat      `json:"io_counters"`
		CPUPercent    float64                  `json:"cpu_percent"`
		CreateTime    int64                    `json:"create_time"`
		MemoryInfo    *pro.MemoryInfoStat      `json:"memory_info"`
		MemoryPercent float32                  `json:"memory_percent"`
		NumThreads    int32                    `json:"num_threads"`
		Nice          int32                    `json:"nice"`
		IONice        int32                    `json:"io_nice"`
	}
)

func GetOSProcessInfo(pid int32) (*OSProcessInfo, error) {
	processInfo := new(OSProcessInfo)
	process, err := pro.NewProcess(pid)
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.Cmdline, err = process.Cmdline()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.Status, err = process.Status()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.Connections, err = process.Connections()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.Times, err = process.Times()
	if IsOSError(err) {
		return processInfo, err
	}

	// Not implemented on darwin, windows
	processInfo.NetIOCounters, err = process.NetIOCounters(true)
	if IsOSError(err) {
		return processInfo, err
	}
	// Not implemented on darwin, windows
	processInfo.Threads, err = process.Threads()
	if IsOSError(err) {
		return processInfo, err
	}

	// Not implemented on darwin, windows
	processInfo.IOCounters, err = process.IOCounters()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.CPUPercent, err = process.CPUPercent()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.CreateTime, err = process.CreateTime()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.MemoryInfo, err = process.MemoryInfo()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.MemoryPercent, err = process.MemoryPercent()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.NumThreads, err = process.NumThreads()
	if IsOSError(err) {
		return processInfo, err
	}

	processInfo.Nice, err = process.Nice()
	if IsOSError(err) {
		return processInfo, err
	}

	// Not implemented on darwin, windows
	processInfo.IONice, err = process.IOnice()
	if IsOSError(err) {
		return processInfo, err
	}

	return processInfo, nil
}
