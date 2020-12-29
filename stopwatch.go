package v60

import (
	"fmt"
	"io"
	"time"
)

type RealStopwatch struct {
	start time.Time
	out   io.Writer
}

func NewRealStopwatch(out io.Writer) *RealStopwatch {
	return &RealStopwatch{start: time.Now(), out: out}
}

func (s RealStopwatch) Wait(duration time.Duration) {
	elapsed := time.Now().Add(duration).Sub(s.start).Round(time.Second)
	fmt.Fprintln(s.out)
	fmt.Fprintf(s.out, "Wait %s for next prompt (elapsed %s)\n", duration, elapsed)
	time.Sleep(duration)
	fmt.Fprintln(s.out)
}
