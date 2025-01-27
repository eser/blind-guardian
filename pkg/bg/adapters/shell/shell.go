package shell

import (
	"errors"
	"fmt"
	"os/exec"
)

var ErrFailedToExecute = errors.New("failed to execute command")

type Shell struct{}

func NewShell() *Shell {
	return &Shell{}
}

func (f *Shell) Run(command string) ([]byte, error) {
	cmd := exec.Command("sh", "-c", command)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("%w: %w", ErrFailedToExecute, err)
	}

	return output, nil
}
