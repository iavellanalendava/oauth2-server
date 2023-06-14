package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestConstants is created to ensure if any change happens to constants
func TestConstants(t *testing.T) {
	assert.Equal(t, "Bearer", Bearer)
	assert.Equal(t, "read", Read)
}
