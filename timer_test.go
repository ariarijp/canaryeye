package canaryeye

import "testing"

func TestExecCommand(t *testing.T) {
	a := map[string]int{
		"127.0.0.1": 100,
	}

	execCommand(a, []string{"cat"})
}
