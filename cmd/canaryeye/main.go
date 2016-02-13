package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/ariarijp/canaryeye"
	"github.com/hpcloud/tail"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: canaryeye FILENAME COMMAND")
		os.Exit(1)
	}

	fn, cmd := os.Args[1], os.Args[2]

	m := map[string]int{}
	begin := time.Now()

	go canaryeye.Run(canaryeye.GetConfig(), &m, &begin, cmd)

	t, err := tail.TailFile(fn, *canaryeye.GetTailConfig())
	canaryeye.HandleError(err)

	r, _ := regexp.Compile("^([^ ]+)")

	for line := range t.Lines {
		matches := r.FindAllStringSubmatch(line.Text, -1)

		if len(matches) > 0 {
			m[matches[0][1]]++
		}
	}
}
