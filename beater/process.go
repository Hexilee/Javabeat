package beater

import (
	"path/filepath"
	"os"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"time"
)

const (
	BeatName = "JavaBeat"
)

type (
	Process struct {
		Pid      int32
		Username string
		OSInfo   *OSProcessInfo
	}
)

func NewProcess(Pid int32, Username string) *Process {
	return &Process{Pid: Pid, Username: Username}
}

func (process *Process) InitOSInfo() *Process {
	osInfo, err := GetOSProcessInfo(process.Pid)
	if err == nil {
		process.OSInfo = osInfo
	}
	return process
}

func (process *Process) GenEvent(b *beat.Beat) beat.Event {
	return beat.Event{
		Timestamp: time.Now(),
		Fields: common.MapStr{
			"type": b.Info.Name,
			"pid":  process.Pid,
			"username": process.Username,
			"os_info": process.InitOSInfo().OSInfo,
		},
	}
}

func (process *Process) Publish(bt *Javabeat, b *beat.Beat) {
	bt.client.Publish(process.GenEvent(b))
}

func GetProcesses() ([]*Process, error) {
	hsperfDirGlob := filepath.Join(os.TempDir(), "hsperfdata_*")
	Processes := make([]*Process, 0)
	logp.Debug(BeatName, "Look for hsperfdata directory, matching glob : %v", hsperfDirGlob)
	hsperfDirs, err := filepath.Glob(hsperfDirGlob)
	if err != nil {
		return nil, err
	}
	for _, dirName := range hsperfDirs {
		dir, err := os.Stat(dirName) // make sure we can walk the dir
		if err != nil {
			logp.Warn("Could not walk dir: %s, skipping it\nError: %s", dir.Name(), err.Error())
			continue
		}
		if dir.IsDir() {
			user, _ := GetUserFromHSperfDir(dir.Name())
			err = filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
				if err == nil && !info.IsDir() {
					if pid, err := AToI32(info.Name()); err == nil {
						Processes = append(Processes, NewProcess(pid, user))
					}
				}
				return nil
			})
		}
	}

	return Processes, err
}

func PublishEvents(bt *Javabeat, b *beat.Beat) {
	processes, err := GetProcesses()
	if err != nil {
		logp.Warn("Get Processes err: %s", err.Error())
		return
	}

	for _, process := range processes {
		go process.Publish(bt, b)
	}
}
