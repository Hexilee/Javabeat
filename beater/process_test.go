package beater

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetProcesses(t *testing.T) {
	processes, err := GetProcesses()
	assert.Nil(t, err)
	for _, process := range processes {
		t.Logf("PID: %s", process.Pid)
		t.Logf("User: %s", process.Username)
	}
}
