package command

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func killCommand(args []string) (string, error) {
	if len(args) == 0 {
		return "", errors.New("kill: enter PID to kill")
	}

	for _, arg := range args {
		pid, err := strconv.Atoi(arg)
		if err != nil {
			return "", fmt.Errorf("kill: wrong PID (%s): %w", arg, err)
		}
		process, err := os.FindProcess(pid)
		if err != nil {
			return "", fmt.Errorf("kill: couln't find process with PID=%v: %w", pid, err)
		}
		if err = process.Kill(); err != nil {
			return "", fmt.Errorf("kill: couln't kill process with PID=%v: %w", pid, err)
		}
	}

	return "", nil
}
