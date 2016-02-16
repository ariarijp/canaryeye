package canaryeye

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTailConfig(t *testing.T) {
	c := GetTailConfig()

	assert.Equal(t, c.Follow, true)
	assert.Equal(t, c.ReOpen, true)
}

func TestHandleError(t *testing.T) {
	HandleError(nil)
}
