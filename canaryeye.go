package canaryeye

import (
	"log"
	"os"

	"github.com/hpcloud/tail"
)

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

func HandleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
