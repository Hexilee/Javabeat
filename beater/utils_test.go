package beater

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetUserFromHSperfDir(t *testing.T) {
	testCases := map[string][]interface{} {
		"hsperfdata_lichenxi": {"lichenxi", true},
		"hsperfdat_lichenxi": {"", false},
		"hsperfdata-lichenxi": {"", false},
	}

	for dirName, results := range testCases {
		expectUser := results[0].(string)
		expectOK := results[1].(bool)
		user, ok := GetUserFromHSperfDir(dirName)
		assert.Equal(t, expectUser, user)
		assert.Equal(t, expectOK, ok)
	}
}
