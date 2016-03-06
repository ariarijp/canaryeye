package canaryeye

import (
	"fmt"
	"os"
	"strings"
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

func TestExecCommand(t *testing.T) {
	a := map[string]int{
		"127.0.0.1": 100,
	}

	execCommand(a, []string{"cat"})
}

func TestGetTailConfig(t *testing.T) {
	c := GetTailConfig()

	assert.Equal(t, c.Follow, true)
	assert.Equal(t, c.ReOpen, true)
}

func TestGetResultSlice(t *testing.T) {
	r := strings.NewReader("{\"results\":[{\"host\":\"127.0.0.1\",\"count\":123}]}")
	res := GetResultSlice(r)

	fmt.Print(res)

	assert.Equal(t, "127.0.0.1", res.Results[0].Host)
	assert.Equal(t, 123, res.Results[0].Count)
}

func TestHandleError(t *testing.T) {
	HandleError(nil)
}
