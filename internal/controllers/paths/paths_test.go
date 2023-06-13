package paths

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestPaths is created to ensure if any change happens to paths
func TestPaths(t *testing.T) {
	assert.Equal(t, "/token", IssueToken)
	assert.Equal(t, "/keys", ListSigningKeys)
	assert.Equal(t, "/verify", VerifyToken)
}
