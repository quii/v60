package v60

import (
	"fmt"
	"io"
	"time"
)

type mockStopwatch struct {
	out     io.Writer
	elapsed time.Duration
}

func newMockStopwatch(out io.Writer) *mockStopwatch {
	return &mockStopwatch{out: out}
}

func (m *mockStopwatch) Wait(duration time.Duration) {
	m.elapsed += duration
	fmt.Fprintf(m.out, "Wait %s for next prompt (elapsed %s)\n", duration, m.elapsed)
}
