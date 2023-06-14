package paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPaths is created to ensure if any change happens to paths
func TestPaths(t *testing.T) {
	assert.Equal(t, "/token", TokenGenerate)
	assert.Equal(t, "/keys", KeysList)
	assert.Equal(t, "/verify", TokenVerify)
}
