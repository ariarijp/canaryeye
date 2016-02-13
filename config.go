package canaryeye

import (
	"errors"
	"os"
	"strconv"
)

type config struct {
	Threshold int
	Sleep     int
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
