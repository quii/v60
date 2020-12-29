package v60

import (
	"fmt"
	"io"
	"time"
)

type RealStopwatch struct {
	out io.Writer
}

func NewRealStopwatch(out io.Writer) *RealStopwatch {
	return &RealStopwatch{out: out}
}

func (s RealStopwatch) Wait(duration time.Duration) {
	fmt.Fprintf(s.out, "\nWait %d seconds\n", int(duration.Seconds()))
	s.scheduleCountdown(duration)
	time.Sleep(duration)
}

func (s RealStopwatch) scheduleCountdown(duration time.Duration) *time.Timer {
	return time.AfterFunc(duration-6*time.Second, func() {
		for i := 5; i > 0; i-- {
			fmt.Fprintln(s.out, i)
		}
	})
}
