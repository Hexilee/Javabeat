package beater

import (
	"path/filepath"
	"os"
	"github.com/elastic/beats/libbeat/logp"
)

const (
	BeatName = "JavaBeat"
)

type (
	Process struct {
		Pid      string
		Username string
	}
)

func NewProcess(Pid, Username string) *Process {
	return &Process{Pid, Username}
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
					Processes = append(Processes, NewProcess(info.Name(), user))
				}
				return nil
			})
		}
	}

	return Processes, err
}

func PushJVMEvent() {

}
