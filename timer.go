package canaryeye

import (
	"encoding/json"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func execCommand(result map[string]int, _cmd []string) {
	b, err := json.Marshal(result)
	HandleError(err)

	var cmd *exec.Cmd
	if len(_cmd) == 0 {
		cmd = exec.Command(_cmd[0])
	} else {
		cmd = exec.Command(_cmd[0], _cmd[1:]...)
	}

	stdin, err := cmd.StdinPipe()
	HandleError(err)

	io.WriteString(stdin, string(b))

	stdin.Close()

	out, err := cmd.CombinedOutput()
	HandleError(err)

	fmt.Println(string(out))
}

func Run(c config, m *map[string]int, begin *time.Time, cmd string) {
	duration := time.Duration(c.Sleep) * time.Second

	for {
		result := map[string]int{}

		for host, cnt := range *m {
			if cnt >= c.Threshold {
				result[host] = cnt
			}
		}

		if len(result) > 0 {
			execCommand(result, strings.Split(cmd, " "))
		}

		*begin = time.Now()
		*m = map[string]int{}

		time.Sleep(duration)
	}
}
