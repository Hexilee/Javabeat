// +build linux

package beater

import (
	"github.com/shirou/gopsutil/cpu"
	pro "github.com/shirou/gopsutil/process"
	"encoding/json"
)

type (
	OSProcessInfo struct {
		NumThreads int32                    `json:"num_threads"`
		Threads    map[int32]*cpu.TimesStat `json:"threads"`
	}
)

func GetOSProcessInfo(pid int32) (string, error) {
	process, err := pro.NewProcess(pid)
	if err != nil {
		return "NewProcess", err
	}
	processInfo := new(OSProcessInfo)
	processInfo.NumThreads, err = process.NumThreads()
	if err != nil {
		return "NumThreads", err
	}

	processInfo.Threads, err = process.Threads()
	if err != nil {
		return "Threads", err
	}

	result, err := json.Marshal(processInfo)
	if err != nil {
		return "json.Marshal", err
	}

	return string(result), nil
}
