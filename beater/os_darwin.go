// +build darwin

package beater

import (
	pro "github.com/shirou/gopsutil/process"
	"encoding/json"
)

type (
	OSProcessInfo struct {
		Cmdline    string  `json:"cmdline"`
		// Not Implemented
		// Threads    map[int32]*cpu.TimesStat `json:"threads"`
		CPUPercent float64 `json:"cpu_percent"`
		CreateTime int64   `json:"create_time"`
		NumThreads int32   `json:"num_threads"`
	}
)

func GetOSProcessInfo(pid int32) (string, error) {
	process, err := pro.NewProcess(pid)
	if err != nil {
		return "NewProcess", err
	}
	processInfo := new(OSProcessInfo)

	processInfo.Cmdline, err = process.Cmdline()
	if err != nil {
		return "Cmdline", err
	}

	// Not implemented
	//processInfo.Threads, err = process.Threads()
	//if err != nil {
	//	return "Threads", err
	//}

	processInfo.CPUPercent, err = process.CPUPercent()
	if err != nil {
		return "CPUPercent", err
	}

	processInfo.CreateTime, err = process.CreateTime()
	if err != nil {
		return "CreateTime", err
	}

	processInfo.NumThreads, err = process.NumThreads()
	if err != nil {
		return "NumThreads", err
	}

	result, err := json.Marshal(processInfo)
	if err != nil {
		return "json.Marshal", err
	}

	return string(result), nil
}
