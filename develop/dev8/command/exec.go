package command

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func execCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("exec: enter command to execute")
	}

	cmd, err := exec.LookPath(args[0])
	if err != nil {
		return "", fmt.Errorf("exec: couldn't find path for %v: %w", args[0], err)
	}
	if cmd == "" {
		return "", fmt.Errorf("exec: couldn't find path for %v", args[0])
	}
	args[0] = cmd

	if err := syscall.Exec(args[0], args, os.Environ()); err != nil {
		return "", fmt.Errorf("exec: %w", err)
	}
	return "", nil
}
