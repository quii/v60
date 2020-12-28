package v60

import (
	"time"
)

type dummyStopwatch struct {
}

func (d dummyStopwatch) Wait(duration time.Duration) {
}

