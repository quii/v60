package v60

import (
	"fmt"
	"io"
	"os/exec"
)

type SayItOutLoud struct {
	textOutput io.Writer
}

func NewSayItOutLoud(out io.Writer) *SayItOutLoud {
	return &SayItOutLoud{textOutput: out}
}

func (s *SayItOutLoud) Write(p []byte) (int, error) {
	cmd := exec.Command("say", string(p))
	err := cmd.Run()
	if err != nil {
		return 0, err
	}
	return fmt.Fprint(s.textOutput, string(p))
}
