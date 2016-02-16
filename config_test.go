package canaryeye

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	os.Setenv("CANARYEYE_SLEEP", "10")
	os.Setenv("CANARYEYE_THRESHOLD", "100")
	c := GetConfig()

	assert.Equal(t, c.Sleep, 10)
	assert.Equal(t, c.Threshold, 100)
}
