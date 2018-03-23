package beater

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetOSProcessInfo(t *testing.T) {
	result, err := GetOSProcessInfo(41413)
	assert.Nil(t, err)
	t.Logf(result)
}
