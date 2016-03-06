package canaryeye

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/hpcloud/tail"
)

type config struct {
	Threshold int
	Sleep     int
}

type ResultSlice struct {
	Results []Result `json:"results"`
}

type Result struct {
	Host  string `json:"host"`
	Count int    `json:"count"`
}

func execCommand(_result map[string]int, _cmd []string) {
	var r ResultSlice

	for k, v := range _result {
		r.Results = append(r.Results, Result{Host: k, Count: v})
	}

	b, err := json.Marshal(r)
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

func GetConfig() config {
	if len(os.Getenv("CANARYEYE_THRESHOLD")) == 0 {
		err := errors.New("CANARYEYE_THRESHOLD env var must be set")
		HandleError(err)
	}

	threshold, err := strconv.Atoi(os.Getenv("CANARYEYE_THRESHOLD"))
	HandleError(err)

	if len(os.Getenv("CANARYEYE_SLEEP")) == 0 {
		err := errors.New("CANARYEYE_SLEEP env var must be set")
		HandleError(err)
	}

	sleep, err := strconv.Atoi(os.Getenv("CANARYEYE_SLEEP"))
	HandleError(err)

	return config{
		Threshold: threshold,
		Sleep:     sleep,
	}
}

func GetTailConfig() *tail.Config {
	return &tail.Config{
		Follow: true,
		ReOpen: true,
		Logger: tail.DiscardingLogger,
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: os.SEEK_END,
		},
	}
}

func GetResultSlice(r io.Reader) ResultSlice {
	jsonBytes, err := ioutil.ReadAll(r)
	HandleError(err)

	var res ResultSlice
	json.Unmarshal(jsonBytes, &res)

	return res
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

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
